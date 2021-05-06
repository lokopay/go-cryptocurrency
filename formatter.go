package cryptocurrency

type Unit struct {
	Unit     string
	Fraction int
	Template string
}

type Formatter struct {
	Code  string
	Units map[string]Unit
}

var formatters = map[string]*Formatter{
	BTC: {
		Code: BTC,
		Units: map[string]Unit{
			BTC: {
				Unit:     "Bitcoin(BTC)",
				Fraction: 1,
				Template: "$unit $value",
			},
		},
	},
}
