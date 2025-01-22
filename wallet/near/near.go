package near

import (
	"crypto/ed25519"
	"fmt"

	"github.com/btcsuite/btcutil/base58"

	"blocchain-wallet/wallet"
	"time"
)

type Wallet struct {
}

func NewWallet() wallet.Wallet {
	return &Wallet{}
}

func (w *Wallet) GenerateAddress() (address string, privateKey string, err error) {
	_, privateKeyByte, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}

	address = fmt.Sprintf("address: %d.testnet", time.Now().UnixMilli())
	//publicKey := hex.EncodeToString(publicKeyByte)
	privateKey = base58.Encode(privateKeyByte)
	return
}
