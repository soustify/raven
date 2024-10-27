package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/soustify/raven/pkg/response"
	"github.com/soustify/raven/pkg/validators"
)

func TransformeResult(c *fiber.Ctx) error {
	statusCode := c.Response().StatusCode()
	body := c.Response().Body()
	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}
	return c.Status(statusCode).JSON(&response.Result{
		Code:    statusCode,
		Content: result,
	})
}

func ValidationMiddleware[T validators.Validatable](c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return err
	}

	statusCode := c.Response().StatusCode()
	body := c.Response().Body()

	if statusCode == fiber.StatusNoContent {
		return nil
	}

	if statusCode == fiber.StatusOK || statusCode == fiber.StatusCreated {
		isArray, err := isJSONArray(string(body))
		if err != nil {
			return err
		} else if isArray {
			_, err = validateAndRespondList[T](body)
			if err != nil {
				return err
			}
		} else {
			_, err = validateAndRespondSingle[T](body)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Valida e responde se o corpo da resposta é uma lista
func validateAndRespondList[T validators.Validatable](body []byte) ([]T, error) {
	var outputList []T
	if err := json.Unmarshal(body, &outputList); err != nil {
		return nil, nil
	}

	for index, item := range outputList {
		if validationErr := item.Validate(); validationErr != nil {
			logrus.WithField("index", index).Errorf("error to validate index: %v error: %v", item, validationErr)
			return nil, fiber.NewError(fiber.StatusInternalServerError, validationErr.Error())
		}
	}

	return outputList, nil
}

func validateAndRespondSingle[T validators.Validatable](body []byte) (*T, error) {
	var output T
	if err := json.Unmarshal(body, &output); err != nil {
		logrus.Errorf("error to decode json: %v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Erro ao processar a resposta")
	}
	if validationErr := output.Validate(); validationErr != nil {
		logrus.Errorf("error to validate error: %v", validationErr)
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
