package cryptocurrency

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	c := New(0.00001, BTC)

	if c.Amount() != 0.00001 {
		t.Errorf("Expected %g got %g", 0.00001, c.Amount())
	}

	if c.Volume() != 1000 {
		t.Errorf("Expected %d got %d", 1000, c.Volume())
	}

	c = NewFromVolume(1234500000, ETH)

	if c.amount.val != 0.000000001234500000 {
		t.Errorf("Expected %g got %g", 0.0000000012345, c.amount.val)
	}

	if c.volume.val != 1234500000 {
		t.Errorf("Expected %d got %d", 1234500000, c.volume.val)
	}

	c = NewFromVolume(-12345, BTC)

	if c.amount.val != -0.00012345 {
		t.Errorf("Expected %g got %g", 0.00012345, c.amount.val)
	}

	if c.volume.val != -12345 {
		t.Errorf("Expected %d got %d", -12345, c.volume.val)
	}

	c = New(0, BTC)

	if c.amount.val != 0 {
		t.Errorf("Expected %d got %g", 0, c.amount.val)
	}

	if c.volume.val != 0 {
		t.Errorf("Expected %d got %d", 0, c.volume.val)
	}
}

func TestCrypto_Allocate(t *testing.T) {
	tcs := []struct {
		volume   int64
		currency string
		ratios   []int
		expect   []int64
	}{
		{974400, BTC, []int{985, 15}, []int64{959784, 14616}},
	}

	for _, tc := range tcs {
		c := NewFromVolume(tc.volume, tc.currency)
		var rs []int64

		split, _ := c.Allocate(tc.ratios...)

		for _, party := range split {
			rs = append(rs, party.volume.val)
		}

		if !reflect.DeepEqual(tc.expect, rs) {
			t.Errorf("Expected allocation of %d for ratios %v to be %v got %v", tc.volume, tc.ratios,
				tc.expect, rs)
		}
	}
}
