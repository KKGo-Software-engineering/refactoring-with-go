package rental

import "fmt"

type Customer struct {
	name    string
	rentals []*Rental
}

func NewCustomer(name string) (rcvr *Customer) {
	rcvr = &Customer{}
	rcvr.rentals = make([]*Rental, 0)
	rcvr.name = name
	return
}
func (rcvr *Customer) AddRental(arg *Rental) {
	rcvr.rentals = append(rcvr.rentals, arg)
}
func (rcvr *Customer) GetName() string {
	return rcvr.name
}
func (rcvr *Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("%v%v%v", "Rental Record for ", rcvr.GetName(), "\n")
	for _, each := range rcvr.rentals {
		thisAmount := 0.0
		switch each.GetMovie().GetPriceCode() {
		case REGULAR:
			thisAmount += 2
			if each.GetDaysRented() > 2 {
				thisAmount += float64(each.GetDaysRented()-2) * 1.5
			}
		case NEW_RELEASE:
			thisAmount += float64(each.GetDaysRented()) * 3.0
		case CHILDRENS:
			thisAmount += 1.5
			if each.GetDaysRented() > 3 {
				thisAmount += float64(each.GetDaysRented()-3) * 1.5
			}
		}
		frequentRenterPoints++
		if each.GetMovie().GetPriceCode() == NEW_RELEASE && each.GetDaysRented() > 1 {
			frequentRenterPoints++
		}
		result += fmt.Sprintf("%v%v%v%.1f%v", "\t", each.GetMovie().GetTitle(), "\t", thisAmount, "\n")
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("%v%.1f%v", "Amount owed is ", totalAmount, "\n")
	result += fmt.Sprintf("%v%v%v", "You earned ", frequentRenterPoints, " frequent renter points")
	return result
}
