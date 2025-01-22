package ethereum

import (
	"blocchain-wallet/wallet"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
}

func NewWallet() wallet.Wallet {
	return &Wallet{}
}

func (w *Wallet) GenerateAddress() (address string, privateKey string, err error) {
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

	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return
}
