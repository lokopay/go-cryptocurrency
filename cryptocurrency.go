package cryptocurrency

import "errors"

// Amount is a datastructure that stores the volume
type Amount struct {
	val int64
}

type Crypto struct {
	amount   *Amount
	currency *Currency
}

func New(amount int64, code string) *Crypto {

	amt := &Amount{val: amount}
	cur := newCurrency(code).get()

	return &Crypto{
		amount:   amt,
		currency: cur,
	}
}

func (c *Crypto) Amount() int64 {
	return c.amount.val
}

func (c *Crypto) IsZero() bool {
	return c.amount.val == 0
}

func (c *Crypto) IsPositive() bool {
	return c.amount.val > 0
}

func (c *Crypto) IsNegative() bool {
	return c.amount.val < 0
}

func (c *Crypto) Allocate(rs ...int) ([]*Crypto, error) {
	if len(rs) == 0 {
		return nil, errors.New("no ratios specified")
	}

	var sum int
	for _, r := range rs {
		sum += r
	}

	var total int64
	cs := make([]*Crypto, 0, len(rs))

	for _, r := range rs {
		party := &Crypto{
			amount:   allocate(c.amount, r, sum),
			currency: c.currency,
		}

		cs = append(cs, party)
		total += party.amount.val
	}

	// Calculate leftover value and divide to first parties.
	lo := c.amount.val - total
	sub := int64(1)
	if lo < 0 {
		sub = -sub
	}

	for p := 0; lo != 0; p++ {
		cs[p].amount = add(cs[p].amount, &Amount{sub})
		lo -= sub
	}

	// for _, cc := range cs {
	// 	cc.amount = volToAmt(cc.volume, cc.currency)
	// }

	return cs, nil
}

// AsMajorUnits lets represent Money struct as subunits (float64) in given Currency value
func (c *Crypto) AsMajorUnits() float64 {
	cur := c.currency.get()

	return cur.Formatter("").ToMajorUnits(c.amount.val)
}

func (c *Crypto) AsUnits(unit string) float64 {
	cur := c.currency.get()
	return cur.Formatter(unit).AsUnits(c.amount.val)
}

func (c *Crypto) Display(unit string) string {
	cur := c.currency.get()
	return cur.Formatter(unit).Format(c.amount.val)
}
