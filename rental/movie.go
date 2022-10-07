package rental

const (
	_ = iota
	CHILDRENS
	NEW_RELEASE
	REGULAR
)

type Movie struct {
	title     string
	priceCode int
}

func NewMovie(title string, priceCode int) (rcvr Movie) {
	return Movie{
		title:     title,
		priceCode: priceCode,
	}
}
func (rcvr Movie) PriceCode() int {
	return rcvr.priceCode
}
func (rcvr Movie) Title() string {
	return rcvr.title
}
