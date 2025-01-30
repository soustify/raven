package headers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/soustify/raven/pkg/response"
)

func GetAuthorization(ctx *fiber.Ctx) (string, error) {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return "", response.NewUnauthorizedError(ctx, "Authorization header is required")
	}

	return authHeader, nil
}

func GetBearerToken(ctx *fiber.Ctx) (string, error) {
	authHeader, err := GetAuthorization(ctx)

	if err != nil {
		return "", err
	}

	if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return "", response.NewUnauthorizedError(ctx, "Invalid Authorization header format")
	}

	return authHeader[7:], nil
}

func GetJwtClaims(ctx *fiber.Ctx) (jwt.MapClaims, error) {
	tokenString, err := GetBearerToken(ctx)
	if err != nil {
		return nil, err
	}

	if tokenString == "" {
		return nil, response.NewUnauthorizedError(ctx, "Invalid Authorization header token is not found")
	}

	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("failed to extract claims from token")
	}
}
