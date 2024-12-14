package request

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetPageNumber(context *fiber.Ctx) int64 {
	page := context.Query("pageNumber", "1")
	intValue, err := strconv.Atoi(page)
	if err != nil {
		logrus.Warn(err)
		return 1
	}
	return int64(intValue)

}

func GetPageSize(context *fiber.Ctx) int64 {
	page := context.Query("pageSize", "10")
	intValue, err := strconv.Atoi(page)
	if err != nil {
		logrus.Warn(err)
		return 10
	}
	return int64(intValue)
}
