package medic

import (
	"crypto/ecdsa"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
)

type (
	Wallet struct {
		PublicKey  *ecdsa.PublicKey
		PrivateKey *ecdsa.PrivateKey
	}
)

// Create new Wallet Account instance using PRIVATE_KEY .env param
func NewWallet() (*Wallet, error) {
	privKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}

	publicKey := privKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("can't parse the public key, creation of new wallet instance failed")
	}

	return &Wallet{
		PublicKey:  publicKeyECDSA,
		PrivateKey: privKey,
	}, nil
}
