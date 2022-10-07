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
func (rcvr Customer) AddRental(arg Rental) {
	rcvr.rentals = append(rcvr.rentals, arg)
}
func (rcvr Customer) Name() string {
	return rcvr.name
}

func amountFor(each Rental) float64 {
	thisAmount := 0.0
	switch each.Movie().PriceCode() {
	case REGULAR:
		thisAmount += 2
		if each.DaysRented() > 2 {
			thisAmount += float64(each.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		thisAmount += float64(each.DaysRented()) * 3.0
	case CHILDRENS:
		thisAmount += 1.5
		if each.DaysRented() > 3 {
			thisAmount += float64(each.DaysRented()-3) * 1.5
		}
	}

	return thisAmount
}

func (rcvr Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("Rental Record for %v\n", rcvr.Name())
	for _, each := range rcvr.rentals {
		thisAmount := amountFor(each)
		frequentRenterPoints++
		if each.Movie().PriceCode() == NEW_RELEASE && each.DaysRented() > 1 {
			frequentRenterPoints++
		}
		result += fmt.Sprintf("\t%v\t%.1f\n", each.Movie().Title(), thisAmount)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}
