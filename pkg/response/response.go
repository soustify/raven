package response

import (
	"github.com/gofiber/fiber/v2"
)

type (
	Result struct {
		Code    int         `json:"code"`
		Content interface{} `json:"content"`
	}
)

func NewError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
}

func NewBadRequestError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
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
	return ctx.Status(fiber.StatusNoContent)
}
