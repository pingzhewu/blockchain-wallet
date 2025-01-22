package main

import (
	"blocchain-wallet/wallet/near"
	"fmt"
)

func main() {
	wallet := near.NewWallet()
	addr, wif, _ := wallet.GenerateAddress()
	fmt.Printf("Address: %s\nWIF: %s\n", addr, wif)
}
