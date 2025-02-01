package headers

import (
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/soustify/raven/pkg/response"
	"strconv"
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

func IsAuthenticated(ctx *fiber.Ctx) (bool, string) {
	authUser := ctx.Get("X-authenticated-user")
	if authUser == "" {
		return false, ""
	}
	return true, authUser
}

func GetExpires(ctx *fiber.Ctx) (bool, int64) {
	expires := ctx.Get("X-expires")
	if expires == "" {
		return false, 0
	}
	val, err := strconv.Atoi(expires)
	if err != nil {
		logrus.Warn("error parsing expires")
		return false, 0
	}
	return true, int64(val)
}

func GetUserPool(ctx *fiber.Ctx) (bool, string) {
	userPool := ctx.Get("X-user-pool")
	if userPool == "" {
		return false, ""
	}
	return true, userPool
}

func GetAuthorities(ctx *fiber.Ctx) (bool, string) {
	authorities := ctx.Get("X-hash-authorities")
	if authorities == "" {
		return false, ""
	}
	decodedBytes, err := base64.StdEncoding.DecodeString(authorities)
	if err != nil {
		fmt.Println("Erro ao decodificar:", err)
		return false, ""
	}
	return true, string(decodedBytes)
}
