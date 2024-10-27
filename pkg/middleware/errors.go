package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soustify/raven/pkg/response"
)

func HandlerErrors(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		if e, ok := err.(*fiber.Error); ok {
			return c.Status(e.Code).JSON(&response.Result{
				Code:    e.Code,
				Content: response.StringResult{Message: e.Message},
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(&response.Result{
			Code:    fiber.StatusInternalServerError,
			Content: response.StringResult{Message: "Erro interno do servidor"},
		})
	}
	return nil
}
