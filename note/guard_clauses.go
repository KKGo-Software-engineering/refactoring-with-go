package main

func getPayAmount() int {
	result := 0
	if isDead() {
		result = deadAmount()
	} else {
		if isSeparated() {
			result = separatedAmount()
		} else {
			if isRetired() {
				result = retiredAmount()
			} else {
				result = normalPayAmount()
			}
		}
	}
	return result
}

func getPayAmount() int {
	if isDead() {
		return deadAmount()
	}
	if isSeparated() {
		return separatedAmount()
	}
	if isRetired() {
		return retiredAmount()
	}
	return normalPayAmount()
}
