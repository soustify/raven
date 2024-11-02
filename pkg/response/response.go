package response

import (
	"github.com/gofiber/fiber/v2"
)

type (
	Result struct {
		Code    int         `json:"code"`
		Content interface{} `json:"content"`
	}

	StringResult struct {
		Message string `json:"message"`
	}
)

func NewError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(&StringResult{
		Message: err.Error(),
	})
}

func NewBadRequestError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&StringResult{
		Message: err.Error(),
	})
}

func NewSuccess(ctx *fiber.Ctx, value interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(value)
}

func NewCreated(ctx *fiber.Ctx, value interface{}) error {
	return ctx.Status(fiber.StatusCreated).JSON(value)
}

func NewAccepted(ctx *fiber.Ctx, value interface{}) error {
	return ctx.Status(fiber.StatusAccepted).JSON(value)
}

func NewNoContent(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNoContent).JSON(&Result{
		Code: fiber.StatusNoContent,
	})
}
