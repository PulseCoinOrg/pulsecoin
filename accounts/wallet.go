package accounts

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/gob"
	"fmt"

	"github.com/PulseCoinOrg/pulsecoin/common"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
	Amount     int64 // wallet balance
}

func New(path string) (*Wallet, error) {
	keyPair, err := GenerateKeys()
	if err != nil {
		return nil, err
	}

	pubKey := keyPair.PublicKey
	privKey := keyPair.PrivateKey

	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)

	addr := common.NewAddr(pubKeyBytes)

	wallet := &Wallet{
		PrivateKey: privKey,
		Address:    addr,
	}

	keyPair.PrintPublicKey()

	err = keyPair.StorePrivateKey(path)
	if err != nil {
		return nil, err
	}

	fmt.Println("your private key was stored at > ", path)

	return wallet, nil
}

type EncodableWallet struct {
	PrivateKey []byte
	Address    common.Address
	Amount     int64
}

func (w *Wallet) Bytes() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	encodable := &EncodableWallet{
		PrivateKey: w.PrivateKey.D.Bytes(),
		Address:    w.Address,
		Amount:     w.Amount,
	}

	if err := enc.Encode(&encodable); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func DecodeWallet(data []byte) *EncodableWallet {
	var wallet *EncodableWallet
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(wallet); err != nil {
		panic(err)
	}
	return wallet
}
