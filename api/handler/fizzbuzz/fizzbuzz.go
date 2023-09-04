package fizzbuzz

import (
	"strconv"
	"strings"
)

type Options struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

func GetResult(options Options) string {
	var result []string
	for i := 1; i <= options.Limit; i++ {
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

		if i%15 == 0 {
			item = options.Str1 + options.Str2
		} else if i%options.Int1 == 0 {
			item = options.Str1
		} else if i%options.Int2 == 0 {
			item = options.Str2
		} else {
			item = strconv.Itoa(i)
		}

		result = append(result, item)
	}

	return strings.Join(result, ",")
}
