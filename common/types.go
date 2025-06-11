package common

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

const (
	HashLength    = 32
	AddressLength = 20
)

type Hash [HashLength]byte

func Sha256Hash(data []byte) Hash {
	return sha256.Sum256(data)
}

func (h Hash) Bytes() []byte {
	return h[:]
}

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

type Address [AddressLength]byte

func NewAddr(pubKeyBytes []byte) Address {
	hasher := ripemd160.New()
	_, err := hasher.Write(pubKeyBytes)
	if err != nil {
		panic(err)
	}
	hash := hasher.Sum(nil)
	return Address(hash[:AddressLength])
}

func (a Address) Bytes() []byte {
	return a[:]
}

func (a Address) String() string {
	return hex.EncodeToString(a[:])
}
