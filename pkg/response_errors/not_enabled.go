package response_errors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	response_result "github.com/soustify/raven/pkg/response"
)

var _ Error = (*NotEnabled)(nil)

type NotEnabled struct {
	message string
}

func (i NotEnabled) ResultStatus(ctx *fiber.Ctx) error {
	return response_result.NewUnprocessable(ctx, fmt.Sprintf("data is not enable to update"))
}

func (i NotEnabled) Error() string {
	return i.message
}

func NewIsNotEnabled(message string) *NotEnabled {
	return &NotEnabled{
		message: message,
	}
}
