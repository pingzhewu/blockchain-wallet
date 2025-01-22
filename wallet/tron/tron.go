package tron

import (
	"blocchain-wallet/wallet"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

type Wallet struct {
}

func NewWallet() wallet.Wallet {
	return &Wallet{}
}

func (w *Wallet) GenerateAddress() (addr string, privateKey string, err error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return
	}

	privateKey = hexutil.Encode(crypto.FromECDSA(key))[2:]
	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return
	}

	addr = address.PubkeyToAddress(*publicKeyECDSA).String()
	return
}
