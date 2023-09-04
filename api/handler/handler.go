package handler

import (
	"fizzbuzz_go/api/handler/fizzbuzz"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FizzBuzz(c *fiber.Ctx) error {

	// check every parameter
	int1, err := strconv.Atoi(c.Query("int1"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid int1 parameter")
	}

	int2, err := strconv.Atoi(c.Query("int2"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid int2 parameter")
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid limit parameter")
	}

	if int1 <= 0 || int2 <= 0 || limit <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Parameters must be positive integers")
	}

	str1 := c.Query("str1")
	str2 := c.Query("str2")
	if str1 == "" || str2 == "" {
		return fiber.NewError(fiber.StatusBadRequest, "str1 and str2 parameters must be non-empty strings")
	}

	fizzbuzzOptions := fizzbuzz.Options{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}

	return c.SendString(fizzbuzz.GetResult(fizzbuzzOptions))
}

/*
func Stats(c *fiber.Ctx) error {
	return c.SendString(fizzbuzz.GetStats())
}
*/
