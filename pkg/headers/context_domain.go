package headers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/soustify/raven/pkg/response"
)

func GetString(ctx *fiber.Ctx, key string) (string, error) {
	value := ctx.Get(key)
	if value == "" {
		return "", response.NewNotFound(ctx, fmt.Sprintf("error to get key: %s", key))
	}

	return value, nil
}

func GetUuid(ctx *fiber.Ctx, key string) (uuid.UUID, error) {
	value := ctx.Get(key)
	if value == "" {
		return uuid.Nil, response.NewNotFound(ctx, fmt.Sprintf("error to get key: %s", key))
	}

	return uuid.FromString(value)
}
