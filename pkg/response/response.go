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
	return ctx.Status(fiber.StatusInternalServerError).JSON(&Result{
		Code: fiber.StatusInternalServerError,
		Content: StringResult{
			Message: err.Error(),
		},
	})
}

func NewBadRequestError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&Result{
		Code: fiber.StatusBadRequest,
		Content: StringResult{
			Message: err.Error(),
		},
	})
}

func NewSuccess(ctx *fiber.Ctx, value interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(&Result{
		Code:    fiber.StatusOK,
		Content: value,
	})
}

func NewCreated(ctx *fiber.Ctx, value interface{}) error {
	return ctx.Status(fiber.StatusCreated).JSON(&Result{
		Code:    fiber.StatusCreated,
		Content: value,
	})
}

func NewAccepted(ctx *fiber.Ctx, value interface{}) error {
	return ctx.Status(fiber.StatusAccepted).JSON(&Result{
		Code:    fiber.StatusAccepted,
		Content: value,
	})
}

func NewNoContent(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNoContent).JSON(&Result{
		Code: fiber.StatusNoContent,
	})
}
