package main

import (
	"fizzbuzz_go/api/databases"
	"io"
	"net/http"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func StartFakeServer(t *testing.T) *fiber.App {
	// mock redis data
	mockRedis := miniredis.RunT(t)
	// start redis with redis mock
	databases.StartRedis(mockRedis.Addr())
	app := StartApi()
	return app
}

func TestFizzBuzzWork(t *testing.T) {
	app := StartFakeServer(t)
	req, err := http.NewRequest("GET", "/?int1=3&int2=5&limit=100&str1=fizz&str2=buzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := app.Test(req, -1)
	assert.Nilf(t, err, "No error should occur")
	// Read the response body
	body, err := io.ReadAll(res.Body)
	assert.Nilf(t, err, "No error should occur")
	assert.Equalf(t, 200, res.StatusCode, "Status code should be 200")
	assert.Equalf(t, string(body), `1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz,fizz,22,23,fizz,buzz,26,fizz,28,29,fizzbuzz,31,32,fizz,34,buzz,fizz,37,38,fizz,buzz,41,fizz,43,44,fizzbuzz,46,47,fizz,49,buzz,fizz,52,53,fizz,buzz,56,fizz,58,59,fizzbuzz,61,62,fizz,64,buzz,fizz,67,68,fizz,buzz,71,fizz,73,74,fizzbuzz,76,77,fizz,79,buzz,fizz,82,83,fizz,buzz,86,fizz,88,89,fizzbuzz,91,92,fizz,94,buzz,fizz,97,98,fizz,buzz`, "Process fizzbuzz until 100")
}

func TestFizzBuzzWithBadParameter(t *testing.T) {
	app := StartFakeServer(t)
	req, err := http.NewRequest("GET", "/?int1=bad&int2=5&limit=100&str1=fizz&str2=buzz", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := app.Test(req, -1)
	assert.Nilf(t, err, "No error should occur")
	// Read the response body
	body, err := io.ReadAll(res.Body)
	assert.Nilf(t, err, "No error should occur")
	assert.Equalf(t, 400, res.StatusCode, "Status code should be 400")
	assert.Equalf(t, string(body), `Invalid int1 parameter`, "int1 parameter is invalid, need to be integer")
}

func TestStatWork(t *testing.T) {
	app := StartFakeServer(t)

	// add 5 times bob and lea, so it should be the most frequent request
	for i := 0; i < 4; i++ {
		req, _ := http.NewRequest("GET", "/?int1=3&int2=5&limit=100&str1=fizz&str2=buzz", nil)
		app.Test(req, -1)
	}
	for i := 0; i < 5; i++ {
		req, _ := http.NewRequest("GET", "/?int1=3&int2=9&limit=100&str1=bob&str2=lea", nil)
		app.Test(req, -1)
	}
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := app.Test(req, -1)
	assert.Nilf(t, err, "No error should occur")
	// Read the response body
	body, err := io.ReadAll(res.Body)
	assert.Nilf(t, err, "No error should occur")
	assert.Equalf(t, 200, res.StatusCode, "Status code should be 200")
	assert.Equalf(t, string(body), `{"int1":3,"int2":9,"limit":100,"str1":"bob","str2":"lea","score":5}`, "Stat result")
}

func Test404(t *testing.T) {
	app := StartFakeServer(t)

	req, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := app.Test(req, -1)
	assert.Nilf(t, err, "No error should occur")
	assert.Equalf(t, 404, res.StatusCode, "Status code should be 404")
}
