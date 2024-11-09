package validators

import "github.com/gofiber/fiber/v2"

type ValidationOutputMiddleware[T Validatable] func(c *fiber.Ctx) error

type Validatable interface {
	Validate() error
}

type Toggleable interface {
	Validatable
	IsEnabled() bool
}
