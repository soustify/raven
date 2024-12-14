package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strconv"
)

func getNumberFromQuery(context *fiber.Ctx, name string, defVal int) int {
	property := context.Query(name, "")
	if property == "" {
		return defVal
	}
	intValue, err := strconv.Atoi(property)
	if err != nil {
		logrus.Warn(err)
		return 1
	}
	return intValue

}

func GetPageNumber(context *fiber.Ctx) int {
	return getNumberFromQuery(context, "pageNumber", 1)
}

func GetPageSize(context *fiber.Ctx) int {
	return getNumberFromQuery(context, "pageSize", 10)
}
