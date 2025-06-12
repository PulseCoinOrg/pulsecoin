package coldwallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"

	"github.com/PulseCoinOrg/pulsecoin/common"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
	Amount     int64 // wallet balance
}

func New() (*Wallet, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	pubKey := privKey.PublicKey
	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)

	addr := common.NewAddr(pubKeyBytes)

	wallet := &Wallet{
		PrivateKey: privKey,
		Address:    addr,
	}

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
