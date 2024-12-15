package response_errors

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/soustify/raven/pkg/response"
)

type (
	Error interface {
		Error() string
		ResultStatus(ctx *fiber.Ctx) error
	}

	GenericError struct {
		message string
	}
)

func (g GenericError) Error() string {
	return g.message
}

func (g GenericError) ResultStatus(ctx *fiber.Ctx) error {
	return response.NewError(ctx, g)
}

func NewGenericError(message string) Error {
	return &GenericError{
		message: message,
	}
}

func ServiceErrorsMiddleware(ctx *fiber.Ctx) error {
	err := ctx.Next()
	if err != nil {
		var e Error
		if errors.As(err, &e) {
			return e.ResultStatus(ctx)
		}
	}
	return nil
}
