package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {
	if num == 3 {
		return "Fizz"
	}

	return fmt.Sprintf("%d", num)
}
