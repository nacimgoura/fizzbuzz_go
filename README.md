# FizzBuzz-Go

## Description

The original fizz-buzz consists in writing all numbers from 1 to 100, and just
replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all
multiples of 15 by "fizzbuzz". The output would look like this:
"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".Your goal is
to implement a web server that will expose a REST API endpoint that:

- Accepts five parameters: three integers int1, int2 and limit, and two strings
  str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples
  of int1 are replaced by str1, all multiples of int2 are replaced by str2, all
  multiples of int1 and int2 are replaced by str1str2.The server needs to be:
- Ready for production
- Easy to maintain by other developersBonus: add a statistics endpoint allowing
  users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the
  number of hits for this request

Use go 1.21 and Docker

## Usage

### CMD

Run service

```bash
go run main.go
```

It's possible to run this service with Docker too (image size 2.6Mo)

```bash
docker build --tag fizzbuzz-go . && docker run -d fizzbuzz-go
```

Execute test (with coverage)

```bash
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
```

or

```bash
make test
```

### HTTP

Possible url paths :

| URL                                            | Method |                                                                             Response |
| ---------------------------------------------- | :----: | -----------------------------------------------------------------------------------: |
| `/?int1=3&int2=5&limit=20&str1=fizz&str2=buzz` |  GET   | `1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,17,fizz,19,buzz` (200) |
| `/stats`                                       |  GET   | `{"int1": 3,"int2": 5,"limit": 100,"str1": "fizz","str2": "buzz","score": 12}` (200) |


Made by Goura Nacim