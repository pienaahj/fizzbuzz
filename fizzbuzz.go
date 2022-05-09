package fizzbuzz

import (
	"fmt"
)

func Run(ints []int) []string {
	strings := []string{}
	for _, i := range ints {
		if i%3 == 0 {
			strings = append(strings, "Fizz")
		} else {
			strings = append(strings, fmt.Sprint(i))
		}

	}
	return strings
}
