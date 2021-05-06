package cryptocurrency

import (
	"math"

	"github.com/shopspring/decimal"
)

// type calculator struct{}

func amtToVol(amt *Amount, cur *Currency) *Volume {

	amont := decimal.NewFromFloat(amt.val)
	unitF := math.Pow10(cur.Fraction)
	fract := decimal.NewFromFloat(unitF)
	volme := amont.Mul(fract).IntPart()

	return &Volume{val: volme}
}

func volToAmt(vol *Volume, cur *Currency) *Amount {

	expon := int32(-1 * cur.Fraction)
	volme := decimal.New(vol.val, expon)
	amont, _ := volme.Float64()

	return &Amount{val: amont}

}

func allocate(v *Volume, r, s int) *Volume {
	return &Volume{v.val * int64(r) / int64(s)}
}

func add(a, b *Volume) *Volume {
	return &Volume{a.val + b.val}
}
