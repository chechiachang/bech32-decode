// from "github.com/binance-chain/go-sdk/common/bech32"
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
)

func main() {
	bech := "tbnb1a6pv5gmnsay4a9sr7nvd0mldz29a6kdxye37ce"
	hrp, data, err := bech32.Decode(bech)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(hrp)
	fmt.Println(data)

	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(converted)
	fmt.Println(hex.EncodeToString(converted))
}
