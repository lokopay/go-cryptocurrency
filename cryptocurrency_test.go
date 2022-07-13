package cryptocurrency

import (
	"testing"
)

func TestNew(t *testing.T) {
	c := New(1000, BTC)

	if c.Amount() != 1000 {
		t.Errorf("Expected %d got %d", 1000, c.Amount())
	}

	if c.AsUnits("sat") != 1000 {
		t.Errorf("Expected %g got %g", 1000.0, c.AsUnits("sat"))
	}

	if c.AsUnits("uBTC") != 10.0 {
		t.Errorf("Expected %g got %g", 10.0, c.AsUnits("uBTC"))
	}

	c = New(0, BTC)
	if c.amount.val != 0 {
		t.Errorf("Expected %d got %d", 0, c.amount.val)
	}

	if c.GetMajorUnit() != "BTC" {
		t.Errorf("Major Unit Expected %s, got %s", "BTC", c.GetMajorUnit())
	}

	c = New(1000, "XBT")
	if c.Amount() != 1000 {
		t.Errorf("Expected %d got %d", 1000, c.Amount())
	}

	if c.AsUnits("sat") != 1000 {
		t.Errorf("Expected %g got %g", 1000.0, c.AsUnits("sat"))
	}

	if c.AsUnits("uBTC") != 10.0 {
		t.Errorf("Expected %g got %g", 10.0, c.AsUnits("uBTC"))
	}

	if c.AsUnits("mBTC") != 0.01 {
		t.Errorf("Expected %g got %g", 0.01, c.AsUnits("mBTC"))
	}

	if c.AsMajorUnits() != 0.00001000 {
		t.Errorf("Expected %g got %g", 0.00001000, c.AsMajorUnits())
	}

	c = New(123456, "ETH")
	if c.Amount() != 123456 {
		t.Errorf("Expected %d got %d", 123456, c.Amount())
	}

	if c.AsUnits("wei") != 123456 {
		t.Errorf("AsUnits Expected %g got %g", 123456.0, c.AsUnits("wei"))
	}

	if c.AsUnits("Kwei") != 123.456 {
		t.Errorf("AsUnits Expected %g got %g", 123.456, c.AsUnits("Kwei"))
	}

	if c.AsUnits("Mwei") != 0.123456 {
		t.Errorf("AsUnits Expected %g got %g", 0.123456, c.AsUnits("Mwei"))
	}

	if c.AsUnits("Gwei") != 0.000123456 {
		t.Errorf("AsUnits Expected %g got %g", 0.000123456, c.AsUnits("Gwei"))
	}

	if c.AsUnits("microether") != 0.000000123456 {
		t.Errorf("AsUnits Expected %g got %g", 0.000000123456, c.AsUnits("microether"))
	}

	if c.AsUnits("milliether") != 0.000000000123456 {
		t.Errorf("AsUnits Expected %g got %g", 0.000000000123456, c.AsUnits("milliether"))
	}

	if c.AsMajorUnits() != 0.000000000000123456 {
		t.Errorf("AsMajorUnits Expected %g got %g", 0.000000000000123456, c.AsMajorUnits())
	}

}

func TestFormat(t *testing.T) {
	tcs := []struct {
		amount   int64
		code     string
		unit     string
		expected string
	}{
		{1000, CryptoCurrencyXBT.String(), "sat", "1000 sat"},
		{1000, CryptoCurrencyXBT.String(), "uBTC", "10.00 uBTC"},
		{1000, CryptoCurrencyXBT.String(), "mBTC", "0.01000 mBTC"},
		{1000, CryptoCurrencyXBT.String(), "BTC", "0.00001000 BTC"},
		{1000, CryptoCurrencyXBT.String(), "", "0.00001000 BTC"},
	}

	for _, tc := range tcs {
		c := New(tc.amount, tc.code)
		r := c.Display(tc.unit)

		if r != tc.expected {
			t.Errorf("Expected formmated %d to be %s got %s", tc.amount, tc.expected, c.Display(tc.unit))
		}
	}
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
