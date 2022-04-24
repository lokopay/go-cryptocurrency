package cryptocurrency

// Cryptocurrencies have unofficially used ISO-like codes on various cryptocurrency exchanges, for instance LTC for Litecoin,
// ADA	6	Ada	Currency on the Cardano platform conflicts with ISO 4217, because AD stands for Andorra.
// BCH	8	Bitcoin Cash
// BNB	8	Binance	BNB conflicts with ISO 4217, because BN stands for Brunei Darussalam.
// BSV	8	Bitcoin SV	BSV (Bitcoin Satoshi Vision) conflicts with ISO 4217, because BS stands for Bahamas.
// BTC, XBT	8	Bitcoin	BTC conflicts with ISO 4217, because BT stands for Bhutan.
// DASH	8	Dash	DASH is of non-standard length.
// DOGE	4	Dogecoin
// EOS	4	EOS
// ETH	18	Ethereum	ETH conflicts with ISO 4217, because ET stands for Ethiopia.
// LTC	8	Litecoin	LTC conflicts with ISO 4217, because LT stands for Lithuania.
// VTC	8	Vertcoin
// XLM	8	Stellar Lumen
// XMR	12	Monero
// XNO[36]	30	Nano
// XRP	6	Ripple
// XTZ	6	Tez
// ZEC	8	Zcash

// type Code string

// const (
// 	BCH  Code = "BCH"
// 	BNB  Code = "BNB"
// 	BTC  Code = "BTC"
// 	DOGE Code = "DOGE"
// 	ETH  Code = "ETH"
// 	LTC  Code = "LTC"
// 	SHIB Code = "SHIB"
// 	XMR  Code = "XMR"
// 	ZEC  Code = "ZEC"
// )
const (
	BCH  = "BCH"
	BNB  = "BNB"
	BTC  = "BTC"
	DOGE = "DOGE"
	ETH  = "ETH"
	LTC  = "LTC"
	SHIB = "SHIB"
	XMR  = "XMR"
	ZEC  = "ZEC"
)

const (
	// bitcoin units
	cBTC = "cBTC"
	mBTC = "mBTC"
	uBTC = "uBTC"
	SAT  = "satoshi"

	// ether units
	wei        = "wei"
	Kwei       = "Kwei"
	Mwei       = "Mwei"
	Gwei       = "Gwei"
	microether = "microether"
	milliether = "milliether"
	ether      = "ether"

	// litecoin units
	mLTC = "mLTC"
	uLTC = "uLTC"
)
