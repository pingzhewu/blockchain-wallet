package wallet

type Wallet interface {
	GenerateAddress() (address string, privateKey string, err error)
}
