package main

func totalAmount(prices []int) int {
	if len(prices) != 0 {
		total := 0
		for i := 0; i < len(prices); i++ {
			total += prices[i]
		}
		return total
	}
	return 0
}

func totalAmount(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	total := 0
	for i := 0; i < len(prices); i++ {
		total += prices[i]
	}

	return total
}
