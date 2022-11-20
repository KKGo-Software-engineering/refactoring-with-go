package main

func getPrice(person Person) string {
	if isChildren(person) {
		return "free"
	}
	return "$15"
}

func isChildren(person Person) string {
	return person.age < 12
}

func getPrice(person Person) string {
	if person.age < 12 {
		return "free"
	}
	return "$15"
}
