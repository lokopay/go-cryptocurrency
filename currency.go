package cryptocurrency

import "strings"

type Currency struct {
	Code     string
	AltCode  string
	Name     string
	Fraction int
	Grapheme string
	Format   string
}

var currencies = map[string]*Currency{
	BTC: {Code: BTC, AltCode: "XBT", Name: "Bitcoin", Fraction: 8, Grapheme: "\u20BF"},
	ETH: {Code: ETH, AltCode: "", Name: "Ether", Fraction: 18, Grapheme: ""},
	LTC: {Code: LTC, AltCode: "", Name: "Litecoin", Fraction: 8, Grapheme: "U+0141"},
}

func newCurrency(code string) *Currency {
	return &Currency{Code: strings.ToUpper(code)}
}

func (c *Currency) getDefault() *Currency {
	return &Currency{Code: c.Code, AltCode: "", Name: "", Fraction: 8, Grapheme: c.Code}
}

func (c *Currency) get() *Currency {
	if curr, ok := currencies[c.Code]; ok {
		return curr
	}

	return c.getDefault()
}

// func (c *Currency) Formatter() *Formatter {

// }
