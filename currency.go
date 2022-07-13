package cryptocurrency

import "strings"

type Currency struct {
	Currency CryptoCurrency
	AltCode  string
	Name     string
	Decimal  string
	Template string
	Major    string
	Units    map[string]Unit
}

type Unit struct {
	Fraction int
	Grapheme string
}

var units = map[CryptoCurrency]map[string]Unit{
	CryptoCurrencyXBT: {
		sat:  {Fraction: 0, Grapheme: "sat"},
		uBTC: {Fraction: 2, Grapheme: "uBTC"},
		mBTC: {Fraction: 5, Grapheme: "mBTC"},
		BTC:  {Fraction: 8, Grapheme: "BTC"},
	},
	CryptoCurrencyETH: {
		wei:        {Fraction: 0, Grapheme: "wei"},
		Kwei:       {Fraction: 3, Grapheme: "Kwei"},
		Mwei:       {Fraction: 6, Grapheme: "Mwei"},
		Gwei:       {Fraction: 9, Grapheme: "Mwei"},
		microether: {Fraction: 12, Grapheme: "microether"},
		milliether: {Fraction: 15, Grapheme: "milliether"},
		ether:      {Fraction: 18, Grapheme: "ether"},
	},
	CryptoCurrencyLTC: {},
}

var currencies = map[CryptoCurrency]*Currency{
	CryptoCurrencyXBT: {Currency: CryptoCurrencyXBT, AltCode: "BTC", Name: "Bitcoin", Decimal: ".", Template: "1 $", Major: "BTC", Units: units[CryptoCurrencyXBT]},
	CryptoCurrencyBTC: {Currency: CryptoCurrencyXBT, AltCode: "BTC", Name: "Bitcoin", Decimal: ".", Template: "1 $", Major: "BTC", Units: units[CryptoCurrencyXBT]},
	CryptoCurrencyETH: {Currency: CryptoCurrencyETH, AltCode: "", Name: "Ether", Decimal: ".", Template: "1 $", Major: "ether", Units: units[CryptoCurrencyETH]},
	CryptoCurrencyLTC: {Currency: CryptoCurrencyLTC, AltCode: "", Name: "Litecoin", Decimal: ".", Template: "1 $", Major: "LTC", Units: units[CryptoCurrencyLTC]},
}

func newCurrency(currency string) *Currency {
	return &Currency{Currency: CryptoCurrency(strings.ToUpper(currency))}
}

func (c *Currency) getDefault() *Currency {
	return &Currency{Currency: c.Currency, AltCode: "BTX", Name: "Bitcoin"}
}

func (c *Currency) Formatter(unit string) *Formatter {
	var u Unit

	if unit == "" {
		u = c.Units[c.Major]
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
	if curr, ok := currencies[c.Currency]; ok {
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
