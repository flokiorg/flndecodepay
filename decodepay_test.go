package decodepay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodepay(t *testing.T) {
	t.Run("decode valid Flokicoin bolt11 lowercase", func(t *testing.T) {
		// Real Flokicoin mainnet invoice (lnfc prefix)
		bolt11 := "lnfc50n1p5hxs4npp5smus34m4mzd2jk7gepqkx9j89w2ukrjfajrvevwgslklzgt8l4zsdqqcqzvsxqyz5vqrzjqwz8th0q6p25q5wcvxh2s75n960tm3tung3vc7lmmmcdltt98pjs9apyqqqqqqqqqyqqqqlgqqqqqeqpjqrzjqtnr4hly8edgpl5wvcx86ekcc2rezdnq2calx5xpwk92l50qscwteapyqqqqqqqqquqqqqlgqqqqqeqpjqrzjqwz8th0q6p25q5wcvxh2s75n960tm3tung3vc7lmmmcdltt98pjs9apyqqqqqqqqqgqqqqlgqqqqqeqpjqsp5ge3wl84an6d34x6r7hm82thlg7vgwdffec5xnut90uwj3fvtaqjs9qxpqysgq6rp4rv2lgkffvd239dmj4ehg7vfxuuulu2amjpjf3rrhjhyay0djvjj8hn95rlz9g2kudqmqhw49urrxm6fhnzx6htwpjv6hj3k53hcq7qxkwk"

		expected := Bolt11{
			Currency:           "fc",
			CreatedAt:          1769161395,
			Expiry:             86400,
			Payee:              "03c9a6e09e2e9d1529c87c034891e05ddf32343fad7a9af563b9de6958c95f5990",
			MLoki:              5000,
			Description:        "",
			DescriptionHash:    "",
			PaymentHash:        "86f908d775d89aa95bc8c8416316472b95cb0e49ec86ccb1c887edf12167fd45",
			MinFinalCLTVExpiry: 400,
		}

		actual, err := Decodepay(bolt11)
		assert.NoError(t, err)
		assert.Equal(t, expected.Currency, actual.Currency)
		assert.Equal(t, expected.CreatedAt, actual.CreatedAt)
		assert.Equal(t, expected.Expiry, actual.Expiry)
		assert.Equal(t, expected.Payee, actual.Payee)
		assert.Equal(t, expected.MLoki, actual.MLoki)
		assert.Equal(t, expected.PaymentHash, actual.PaymentHash)
		assert.Equal(t, expected.MinFinalCLTVExpiry, actual.MinFinalCLTVExpiry)
		assert.NotNil(t, actual.Route)
		assert.Len(t, actual.Route, 3) // Has 3 route hints
	})

	t.Run("decode valid Flokicoin bolt11 uppercase", func(t *testing.T) {
		// Real Flokicoin mainnet invoice (LNFC prefix, uppercase)
		bolt11 := "LNFC50N1P5HXS4NPP5SMUS34M4MZD2JK7GEPQKX9J89W2UKRJFAJRVEVWGSLKLZGT8L4ZSDQQCQZVSXQYZ5VQRZJQWZ8TH0Q6P25Q5WCVXH2S75N960TM3TUNG3VC7LMMMCDLTT98PJS9APYQQQQQQQQQYQQQQLGQQQQQEQPJQRZJQTNR4HLY8EDGPL5WVCX86EKCC2REZDNQ2CALX5XPWK92L50QSCWTEAPYQQQQQQQQQUQQQQLGQQQQQEQPJQRZJQWZ8TH0Q6P25Q5WCVXH2S75N960TM3TUNG3VC7LMMMCDLTT98PJS9APYQQQQQQQQQGQQQQLGQQQQQEQPJQSP5GE3WL84AN6D34X6R7HM82THLG7VGWDFFEC5XNUT90UWJ3FVTAQJS9QXPQYSGQ6RP4RV2LGKFFVD239DMJ4EHG7VFXUUULU2AMJPJF3RRHJHYAY0DJVJJ8HN95RLZ9G2KUDQMQHW49URRXM6FHNZX6HTWPJV6HJ3K53HCQ7QXKWK"

		actual, err := Decodepay(bolt11)
		assert.NoError(t, err)
		assert.Equal(t, "fc", actual.Currency)
		assert.Equal(t, int64(5000), actual.MLoki)
	})

	t.Run("fails to decode mixed case invoice", func(t *testing.T) {
		// Mixed case invoices are invalid per BOLT11 spec
		bolt11 := "LNFC50N1P5HXS4NPP5SMUS34M4MZD2JK7gepqkx9j89w2ukrjfajrvevwgslklzgt8l4zsdqqcqzvsxqyz5vqrzjqwz8th0q6p25q5wcvxh2s75n960tm3tung3vc7lmmmcdltt98pjs9apyqqqqqqqqqyqqqqlgqqqqqeqpjqrzjqtnr4hly8edgpl5wvcx86ekcc2rezdnq2calx5xpwk92l50qscwteapyqqqqqqqqquqqqqlgqqqqqeqpjqrzjqwz8th0q6p25q5wcvxh2s75n960tm3tung3vc7lmmmcdltt98pjs9apyqqqqqqqqqgqqqqlgqqqqqeqpjqsp5ge3wl84an6d34x6r7hm82thlg7vgwdffec5xnut90uwj3fvtaqjs9qxpqysgq6rp4rv2lgkffvd239dmj4ehg7vfxuuulu2amjpjf3rrhjhyay0djvjj8hn95rlz9g2kudqmqhw49urrxm6fhnzx6htwpjv6hj3k53hcq7qxkwk"

		_, err := Decodepay(bolt11)
		assert.Error(t, err)
	})

	t.Run("Returns error for invalid bolt11 invoice", func(t *testing.T) {
		bolt11 := "lnfc1234Invalid"

		_, err := Decodepay(bolt11)
		assert.Error(t, err)
	})
}
