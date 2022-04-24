package cryptocurrency

// type calculator struct{}

// func amtToVol(amt *Amount, cur *Currency) *Amount {

// 	amont := decimal.NewFromFloat(amt.val)
// 	unitF := math.Pow10(cur.Fraction)
// 	fract := decimal.NewFromFloat(unitF)
// 	volme := amont.Mul(fract).IntPart()

// 	return &Volume{val: volme}
// }

// func volToAmt(vol *Amount, cur *Currency) *Amount {

// 	expon := int32(-1 * cur.Fraction)
// 	volme := decimal.New(vol.val, expon)
// 	amont, _ := volme.Float64()

// 	return &Amount{val: amont}

// }

func allocate(v *Amount, r, s int) *Amount {
	return &Amount{v.val * int64(r) / int64(s)}
}

func add(a, b *Amount) *Amount {
	return &Amount{a.val + b.val}
}
