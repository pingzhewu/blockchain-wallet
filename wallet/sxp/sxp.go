package sxp

import (
	"blocchain-wallet/wallet"
	"crypto/sha256"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcutil/base58"
	r160 "golang.org/x/crypto/ripemd160"
)

type Wallet struct {
}

func NewWallet() wallet.Wallet {
	return &Wallet{}
}

func (w *Wallet) GenerateAddress() (address string, privateKey string, err error) {
	key, err := btcec.NewPrivateKey()
	if err != nil {
		return
	}

	result := make([]byte, 34)
	result[0] = 252
	copy(result[1:], key.Serialize())
	result[33] = 0x01

	privateKey = base58.Encode(result)

	// 推出地址
	pubCompress := key.PubKey().SerializeCompressed()
	h := r160.New()
	h.Write(pubCompress)
	hByte := h.Sum(nil)

	payload := make([]byte, 21)

	payload[0] = 30

	copy(payload[1:], hByte)

	// encode
	checksum := sha256.Sum256(payload)
	dataWithChecksum := append(payload, checksum[:4]...)

	address = base58.Encode(dataWithChecksum)

	return
}
