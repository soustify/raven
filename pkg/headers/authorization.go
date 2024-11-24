package headers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/soustify/raven/pkg/response"
)

func GetAuthorization(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return "", response.NewUnauthorizedError(ctx, "Authorization header is required")
	}

	if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return "", response.NewUnauthorizedError(ctx, "Invalid Authorization header format")
	}

	return authHeader[7:], nil
}
