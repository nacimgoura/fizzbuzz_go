package handler

import (
	"context"
	"fizzbuzz_go/api/databases"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type OptionsFizzBuzz struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Score int    `json:"score"`
}

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

	if databases.RedisClient != nil {
		// add value in sort list in redis to know what the most frequent request has been
		key := fmt.Sprintf("%d|%d|%d|%s|%s", int1, int2, limit, str1, str2)
		databases.RedisClient.ZIncrBy(context.Background(), "stats", 1.0, key).Err()
	}

	var result []string
	for i := 1; i <= limit; i++ {
		var item string

		// better way
		/*if i%int1 == 0 {
			item += str1
		}
		if i%int2 == 0 {
			item += str2
		}
		if item == "" {
			item = strconv.Itoa(i)
		}*/

		if i%int1 == 0 && i%int2 == 0 {
			item = str1 + str2
		} else if i%int1 == 0 {
			item = str1
		} else if i%int2 == 0 {
			item = str2
		} else {
			item = strconv.Itoa(i)
		}

		result = append(result, item)
	}

	return c.SendString(strings.Join(result, ","))
}

func Stats(c *fiber.Ctx) error {
	if databases.RedisClient == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to get stats")
	}
	// get the first element of the sorted list (the most frequent request)
	query, err := databases.RedisClient.ZRevRange(context.Background(), "stats", 0, 0).Result()
	if err != nil || query == nil || len(query) == 0 {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to get stats")
	}
	listQuery := strings.Split(query[0], "|")
	score, err := databases.RedisClient.ZScore(context.Background(), "stats", query[0]).Result()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Unable to get stats")
	}
	fzOptions := OptionsFizzBuzz{
		Score: int(score),
		Str1:  listQuery[3],
		Str2:  listQuery[4],
	}
	fzOptions.Int1, _ = strconv.Atoi(listQuery[0])
	fzOptions.Int2, _ = strconv.Atoi(listQuery[1])
	fzOptions.Limit, _ = strconv.Atoi(listQuery[2])
	return c.JSON(fzOptions)
}
