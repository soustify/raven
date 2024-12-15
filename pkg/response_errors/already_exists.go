package response_errors

import (
	"github.com/gofiber/fiber/v2"
	response_result "github.com/soustify/raven/pkg/response"
)

var _ Error = (*AlreadyExists)(nil)

type (
	AlreadyExists struct {
		message string
	}
)

func (a AlreadyExists) ResultStatus(ctx *fiber.Ctx) error {
	return response_result.NewConflict(ctx, a.message)
}

func (a AlreadyExists) Error() string {
	return a.message
}

func NewAlreadyExists(message string) *AlreadyExists {
	return &AlreadyExists{
		message: message,
	}
}
