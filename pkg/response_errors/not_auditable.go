package response_errors

import (
	"github.com/gofiber/fiber/v2"
	response_result "github.com/soustify/raven/pkg/response"
)

var _ Error = (*NotAuditable)(nil)

type NotAuditable struct {
	message string
}

func (n NotAuditable) Error() string {
	return n.message
}

func (n NotAuditable) ResultStatus(ctx *fiber.Ctx) error {
	return response_result.NewError(ctx, n)
}

func NewNotAuditable(message string) *NotAuditable {
	return &NotAuditable{
		message: message,
	}
}
