package main

import "fmt"

func printReceipt(person Person) {
	amount := getPrice(person)

	//receipt details
	fmt.Println("name:", person.Name)
	fmt.Println("total amount:", amount)
}

func printReceipt(person Person) {
	amount := getPrice(person)

	printDetails(person, amount)
}

func printDetails(person, amount) {
	fmt.Println("name:", person.Name)
	fmt.Println("total amount:", amount)
}
