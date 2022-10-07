package rental

const CHILDRENS = 2
const NEW_RELEASE = 1
const REGULAR = 0

type Movie struct {
	title     string
	priceCode int
}

func NewMovie(title string, priceCode int) (rcvr *Movie) {
	rcvr = &Movie{}
	rcvr.title = title
	rcvr.priceCode = priceCode
	return
}
func (rcvr *Movie) GetPriceCode() int {
	return rcvr.priceCode
}
func (rcvr *Movie) GetTitle() string {
	return rcvr.title
}
func (rcvr *Movie) SetPriceCode(arg int) {
	rcvr.priceCode = arg
}
