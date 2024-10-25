package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/soustify/raven/pkg/response"
)

func headerMiddleware(contextFiber *fiber.Ctx, key string) (string, error) {
	value := contextFiber.Get(key)
	if value == "" {
		return "", response.NewBadRequestError(contextFiber, errors.New(fmt.Sprintf("%s is mandatory on header", key)))
	}
	return value, nil
}

func NewHeaderValidator(headers ...string) func(contextFiber *fiber.Ctx) error {
	return func(contextFiber *fiber.Ctx) error {
		var ctxWithContext context.Context

		for _, header := range headers {
			value, err := headerMiddleware(contextFiber, header)
			if err != nil {
				return err
			}
			if ctxWithContext == nil {
				ctxWithContext = context.WithValue(contextFiber.UserContext(), header, value)
			} else {
				ctxWithContext = context.WithValue(ctxWithContext, header, value)
			}
			contextFiber.SetUserContext(ctxWithContext)
		}
		return contextFiber.Next()
	}
}
