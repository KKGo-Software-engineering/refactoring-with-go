package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) (rcvr Customer) {
	return Customer{
		name:    name,
		rentals: []Rental{},
	}
}
func (c Customer) AddRental(arg Rental) {
	c.rentals = append(c.rentals, arg)
}
func (c Customer) Name() string {
	return c.name
}

func getPoints(r Rental) int {
	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		return 2
	}

	return 1
}

func getTotalPoints(rentals []Rental) int {
	points := 0
	for _, r := range rentals {
		points += getPoints(r)
	}

	return points
}

func getTotalAmount(rentals []Rental) float64 {
	result := 0.0
	for _, r := range rentals {
		result += r.Charge()
	}
	return result
}

func (c Customer) Statement() string {
	totalAmount := getTotalAmount(c.rentals)
	points := getTotalPoints(c.rentals)

	result := fmt.Sprintf("Rental Record for %v\n", c.Name())
	for _, r := range c.rentals {
		title := r.Movie().Title()
		amount := r.Charge()

		result += fmt.Sprintf("\t%v\t%.1f\n", title, amount)
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", points)
	return result
}
