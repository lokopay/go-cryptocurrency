package cryptocurrency

import "strings"

type Currency struct {
	Code     string
	AltCode  string
	Name     string
	Decimal  string
	Template string
	Units    map[string]Unit
}

type Unit struct {
	Fraction int
	Grapheme string
}

var units = map[string]map[string]Unit{
	BTC: {
		BTC:  {Fraction: 8, Grapheme: "BTC"},
		mBTC: {Fraction: 5, Grapheme: "mBTC"},
		uBTC: {Fraction: 3, Grapheme: "uBTC"}},
	ETH: {
		wei:  {Fraction: 18, Grapheme: "wei"},
		Kwei: {Fraction: 18, Grapheme: "wei"},
		Mwei: {Fraction: 18, Grapheme: "wei"},
	},
}

var major = map[string]string{
	BTC: BTC,
}

var currencies = map[string]*Currency{
	BTC: {Code: BTC, AltCode: "XBT", Name: "Bitcoin", Decimal: ".", Template: "1 $", Units: units[BTC]},
	ETH: {Code: ETH, AltCode: "", Name: "Ether", Decimal: ".", Template: "1 $", Units: units[ETH]},
	LTC: {Code: LTC, AltCode: "", Name: "Litecoin", Decimal: ".", Template: "1 $", Units: units[LTC]},
}

func newCurrency(code string) *Currency {
	return &Currency{Code: strings.ToUpper(code)}
}

func (c *Currency) getDefault() *Currency {
	return &Currency{Code: c.Code, AltCode: "BTX", Name: "Bitcoin"}
}

func (c *Currency) Formatter(unit string) *Formatter {
	var u Unit

	if unit == "" {
		u = c.Units[c.Code]
	} else {
		u = c.Units[unit]
	}

	return &Formatter{
		Fraction: u.Fraction,
		Decimal:  c.Decimal,
		Grapheme: u.Grapheme,
		Template: c.Template,
	}
}

func (c *Currency) get() *Currency {
	if curr, ok := currencies[c.Code]; ok {
		return curr
	}

	return c.getDefault()
}

// func (c *Currency) Formatter() *Formatter {

// }

// var formatters = map[string]*Formatter{
// 	BTC: {
// 		Code: BTC,
// 		Units: map[string]Unit{
// 			BTC: {
// 				Unit:     "Bitcoin(BTC)",
// 				Fraction: 8,
// 				Template: "$unit $value",
// 			},
// 			mBTC: {
// 				Unit:     "mBTC",
// 				Fraction: 5,
// 			},
// 		},
// 	},
// }raction: 8, Grapheme: "U+0141"

// var units =
