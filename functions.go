package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Sprintf("%s environment variable not set.", k))
	}
	return v
}

func canAmountToFloat(amount *big.Int) *big.Float {
	famount := new(big.Float)
	famount.SetString(amount.String())
	return new(big.Float).Quo(famount, big.NewFloat(math.Pow10(6)))
}
