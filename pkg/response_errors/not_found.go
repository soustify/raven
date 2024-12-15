package response_errors

import (
	"github.com/gofiber/fiber/v2"
	response_result "github.com/soustify/raven/pkg/response"
)

var _ Error = (*NotFound)(nil)

type NotFound struct {
	message string
}

func (n NotFound) ResultStatus(ctx *fiber.Ctx) error {
	return response_result.NewNotFound(ctx, n.message)
}

func (n NotFound) Error() string {
	return n.message
}

func NewNotFound(message string) *NotFound {
	return &NotFound{
		message: message,
	}
}
