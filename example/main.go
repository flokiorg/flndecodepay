package main

import (
	"encoding/json"
	"fmt"

	decodepay "github.com/flokiorg/flndecodepay"
)

func main() {
	// Decode a Flokicoin Lightning invoice (lnfc prefix)
	bolt11, _ := decodepay.Decodepay("lnfc50n1p5hxs4npp5smus34m4mzd2jk7gepqkx9j89w2ukrjfajrvevwgslklzgt8l4zsdqqcqzvsxqyz5vqrzjqwz8th0q6p25q5wcvxh2s75n960tm3tung3vc7lmmmcdltt98pjs9apyqqqqqqqqqyqqqqlgqqqqqeqpjqrzjqtnr4hly8edgpl5wvcx86ekcc2rezdnq2calx5xpwk92l50qscwteapyqqqqqqqqquqqqqlgqqqqqeqpjqrzjqwz8th0q6p25q5wcvxh2s75n960tm3tung3vc7lmmmcdltt98pjs9apyqqqqqqqqqgqqqqlgqqqqqeqpjqsp5ge3wl84an6d34x6r7hm82thlg7vfxuuulu2amjpjf3rrhjhyay0djvjj8hn95rlz9g2kudqmqhw49urrxm6fhnzx6htwpjv6hj3k53hcq7qxkwkwdffec5xnut90uwj3fvtaqjs9qxpqysgq6rp4rv2lgkffvd239dmj4ehg7vfxuuulu2amjpjf3rrhjhyay0djvjj8hn95rlz9g2kudqmqhw49urrxm6fhnzx6htwpjv6hj3k53hcq7qxkwk")
	j, _ := json.MarshalIndent(bolt11, "", "  ")
	fmt.Println(string(j))
}
