package rental

const (
	_ = iota
	CHILDRENS
	NEW_RELEASE
	REGULAR
)

type Pricer interface {
	Charge(daysRented int) float64
	PriceCode() int
}

type Childrens struct {
	priceCode int
}

func (c Childrens) Charge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}

func (c Childrens) PriceCode() int {
	return c.priceCode
}

func CreateChildrens() Childrens {
	return Childrens{
		priceCode: CHILDRENS,
	}
}

type Regulars struct {
	priceCode int
}

func (r Regulars) Charge(daysRented int) float64 {
	result := 2.0
	if daysRented > 2 {
		result += float64(daysRented-2) * 1.5
	}
	return result
}

func (r Regulars) PriceCode() int {
	return r.priceCode
}

func CreateRegulars() Regulars {
	return Regulars{
		priceCode: REGULAR,
	}
}

type NewRelease struct {
	priceCode int
}

func (n NewRelease) Charge(daysRented int) float64 {
	return float64(daysRented) * 3
}

func (n NewRelease) PriceCode() int {
	return n.priceCode
}

func CreateNewRelease() NewRelease {
	return NewRelease{
		priceCode: NEW_RELEASE,
	}
}

type Movie struct {
	title     string
	priceCode int
	price     Pricer
}

func NewMovie(title string, charger Pricer) Movie {
	return Movie{
		title:     title,
		priceCode: charger.PriceCode(),
		price:     charger,
	}
}

func (m Movie) PriceCode() int {
	return m.priceCode
}
func (m Movie) Title() string {
	return m.title
}

func (m Movie) Charge(daysRented int) float64 {
	return m.price.Charge(daysRented)
}
