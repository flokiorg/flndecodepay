# flndecodepay

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/flokiorg/flndecodepay/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/flokiorg/flndecodepay)

Simple Flokicoin Lightning Network BOLT11 invoice decoder with outputs similar to [c-lightning](https://github.com/ElementsProject/lightning/blob/master/doc/lightning-decodepay.7.txt) using code from [flnd](https://github.com/flokiorg/flnd).

This is a fork of [ln-decodepay](https://github.com/nbd-wtf/ln-decodepay) adapted for the Flokicoin Lightning Network.

### install

```
go get -u github.com/flokiorg/flndecodepay
```

### use

```go
package main

import (
	"encoding/json"
	"fmt"

	decodepay "github.com/flokiorg/flndecodepay"
)

func main() {
	// Decode a Flokicoin Lightning invoice (lnfc prefix)
	bolt11, _ := decodepay.Decodepay("lnfc50n1p5hxs4npp5smus34m4mzd2jk7gepqkx9j89w2ukrjfajrvevwgslklzgt8l4zsdqqcqzvsxqyz5vqrzjqwz8th0q6p25q5wcvxh2s75n960tm3tung3vc7lmmmcdltt98pjs9apyqqqqqqqqqyqqqqlgqqqqqeqpjqrzjqtnr4hly8edgpl5wvcx86ekcc2rezdnq2calx5xpwk92l50qscwteapyqqqqqqqqquqqqqlgqqqqqeqpjqrzjqwz8th0q6p25q5wcvxh2s75n960tm3tung3vc7lmmmcdltt98pjs9apyqqqqqqqqqgqqqqlgqqqqqeqpjqsp5ge3wl84an6d34x6r7hm82thlg7vgwdffec5xnut90uwj3fvtaqjs9qxpqysgq6rp4rv2lgkffvd239dmj4ehg7vfxuuulu2amjpjf3rrhjhyay0djvjj8hn95rlz9g2kudqmqhw49urrxm6fhnzx6htwpjv6hj3k53hcq7qxkwk")
	j, _ := json.MarshalIndent(bolt11, "", "  ")
	fmt.Println(string(j))
}
```

outputs

```json
{
  "currency": "fc",
  "created_at": 1769161395,
  "expiry": 86400,
  "payee": "03c9a6e09e2e9d1529c87c034891e05ddf32343fad7a9af563b9de6958c95f5990",
  "mloki": 5000,
  "payment_hash": "86f908d775d89aa95bc8c8416316472b95cb0e49ec86ccb1c887edf12167fd45",
  "min_final_cltv_expiry": 400,
  "routes": [
    [
      {
        "pubkey": "038475dde0d0554051d861aea87a932e9ebdc57c9a22cc7bfbdef0dfad65386502",
        "short_channel_id": "16000000x0x1",
        "fee_base_mloki": 1000,
        "fee_proportional_millionths": 100,
        "cltv_expiry_delta": 400
      }
    ]
  ]
}
```

### Supported currencies

- `lnfc` - Flokicoin mainnet
- `lnfcrt` - Flokicoin regtest
- `lnfct` - Flokicoin testnet
