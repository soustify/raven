package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/soustify/raven/pkg/response"
	"github.com/soustify/raven/pkg/validators"
)

func ValidationMiddleware[T validators.Validatable](c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return err
	}

	statusCode := c.Response().StatusCode()
	body := c.Response().Body()

	var result interface{}

	if statusCode == fiber.StatusNoContent {
		return defaultResponse(c, "", statusCode)
	}

	if statusCode == fiber.StatusOK || statusCode == fiber.StatusCreated {
		isArray, err := isJSONArray(string(body))
		if err != nil {
			return err
		} else if isArray {
			result, err = validateAndRespondList[T](body)
			if err != nil {
				return err
			}
		} else {
			result, err = validateAndRespondSingle[T](body)
			if err != nil {
				return err
			}
		}
	} else {
		if err := json.Unmarshal(body, &result); err != nil {
			return err
		}
	}

	return defaultResponse(c, result, statusCode)
}

// Valida e responde se o corpo da resposta é uma lista
func validateAndRespondList[T validators.Validatable](body []byte) ([]T, error) {
	var outputList []T
	if err := json.Unmarshal(body, &outputList); err != nil {
		return nil, nil
	}

	for _, item := range outputList {
		if validationErr := item.Validate(); validationErr != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, validationErr.Error())
		}
	}

	return outputList, nil
}

func validateAndRespondSingle[T validators.Validatable](body []byte) (*T, error) {
	var output T
	if err := json.Unmarshal(body, &output); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Erro ao processar a resposta")
	}
	if validationErr := output.Validate(); validationErr != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, validationErr.Error())
	}
	return &output, nil
}

func isJSONArray(jsonStr string) (bool, error) {
	jsonStr = jsonStr[:len(jsonStr)]
	if len(jsonStr) == 0 {
		return false, fmt.Errorf("json string is empty")
	}
	switch jsonStr[0] {
	case '[':
		return true, nil
	case '{':
		return false, nil
	default:
		return false, fmt.Errorf("invalid json format")
	}
}

// Retorna a resposta padrão
func defaultResponse(c *fiber.Ctx, body interface{}, statusCode int) error {
	return c.Status(statusCode).JSON(&response.Result{
		Code:    statusCode,
		Content: body,
	})
}
