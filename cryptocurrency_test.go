package cryptocurrency

import (
	"testing"
)

func TestNew(t *testing.T) {
	c := New(1000, BTC)

	if c.Amount() != 1000 {
		t.Errorf("Expected %d got %d", 1000, c.Amount())
	}

	c = New(0, BTC)

	if c.amount.val != 0 {
		t.Errorf("Expected %d got %d", 0, c.amount.val)
	}

	c = New(10000, "BTC")
	if c.Amount() != 10000 {
		t.Errorf("Expected %d got %d", 10000, c.Amount())
	}

	if c.AsMajorUnits() != 0.00010000 {
		t.Errorf("Expected %g got %g", 0.00010000, c.AsMajorUnits())
	}

	if c.AsUnits("mBTC") != 0.00010000 {
		t.Errorf("Expected %g got %s", 0.00010000, c.Display("mBTC"))
	}

	c.Display("")

}

// func TestCrypto_Allocate(t *testing.T) {
// 	tcs := []struct {
// 		volume   int64
// 		currency string
// 		ratios   []int
// 		expect   []int64
// 	}{
// 		{974400, BTC, []int{985, 15}, []int64{959784, 14616}},
// 	}

// 	for _, tc := range tcs {
// 		c := NewFromVolume(tc.volume, tc.currency)
// 		var rs []int64

// 		split, _ := c.Allocate(tc.ratios...)

// 		for _, party := range split {
// 			rs = append(rs, party.volume.val)
// 		}

// 		if !reflect.DeepEqual(tc.expect, rs) {
// 			t.Errorf("Expected allocation of %d for ratios %v to be %v got %v", tc.volume, tc.ratios,
// 				tc.expect, rs)
// 		}
// 	}
// }
