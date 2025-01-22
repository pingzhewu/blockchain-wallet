package btc

import (
	"blocchain-wallet/wallet"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

type Wallet struct {
}

func NewWallet() wallet.Wallet {
	return &Wallet{}
}

func (w *Wallet) GenerateAddress() (address string, privateKey string, err error) {
	cfg := &chaincfg.MainNetParams
	key, err := btcec.NewPrivateKey()
	if err != nil {
		return
	}

	wif, err := btcutil.NewWIF(key, cfg, true)
	if err != nil {
		return
	}

	privateKey = wif.String()
	taprootAddr, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(
		wif.PrivKey.PubKey()),
		cfg,
	)

	if err != nil {
		return
	}

	address = taprootAddr.String()
	return

}
