package cryptocurrency

import "errors"

// Amount is a datastructure that stores the volume
type Amount struct {
	val float64
}

type Volume struct {
	val int64
}

type Crypto struct {
	amount   *Amount
	volume   *Volume
	currency *Currency
}

func New(amount float64, code string) *Crypto {

	amt := &Amount{val: amount}
	cur := newCurrency(code).get()
	vol := amtToVol(amt, cur)

	return &Crypto{
		amount:   amt,
		volume:   vol,
		currency: cur,
	}
}

func NewFromVolume(volume int64, code string) *Crypto {

	vol := &Volume{val: volume}
	cur := newCurrency(code).get()
	amt := volToAmt(vol, cur)

	return &Crypto{
		amount:   amt,
		volume:   vol,
		currency: cur,
	}
}

func (c *Crypto) Volume() int64 {
	return c.volume.val
}

func (c *Crypto) Amount() float64 {
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
			volume:   allocate(c.volume, r, sum),
			currency: c.currency,
		}

		cs = append(cs, party)
		total += party.volume.val
	}

	// Calculate leftover value and divide to first parties.
	lo := c.volume.val - total
	sub := int64(1)
	if lo < 0 {
		sub = -sub
	}

	for p := 0; lo != 0; p++ {
		cs[p].volume = add(cs[p].volume, &Volume{sub})
		lo -= sub
	}

	for _, cc := range cs {
		cc.amount = volToAmt(cc.volume, cc.currency)
	}

	return cs, nil
}

func (c *Crypto) Display() string {

}
