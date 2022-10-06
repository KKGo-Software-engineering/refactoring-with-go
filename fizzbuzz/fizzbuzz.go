package fizzbuzz

import "fmt"

func FizzBuzz(num int) string {
	if isFizzBuzz(num) {
		return "FizzBuzz"
	}

	if isBuzz(num) {
		return "Buzz"
	}

	if isFizz(num) {
		return "Fizz"
	}

	return fmt.Sprintf("%d", num)
}

func isFizzBuzz(num int) bool {
	return num%15 == 0
}

func isFizz(num int) bool {
	return num%3 == 0
}

func isBuzz(num int) bool {
	return num%5 == 0
}
