package accounts

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrPwdNotSet   = errors.New("password not set, you must set a password")
	ErrWrongPwd    = errors.New("password is incorrect")
)

type KeyGen interface {
	New() error

	GetPublicKeyBytes(privKey *ecdsa.PrivateKey) []byte

	// for the console
	PrintPublicKey()

	StorePrivateKey(path string) error
	ViewPrivateKey(path string, pwd string) error // TODO implement
}

type KeyPair struct {
	KeyGen
	Pwd        string
	SigningKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func NewKeyPair(pwd string) (*KeyPair, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	if pwd == "" {
		return nil, ErrPwdNotSet
	}

	return &KeyPair{
		Pwd:        pwd,
		SigningKey: privKey,
		PublicKey:  &privKey.PublicKey,
	}, nil
}

func (kp *KeyPair) GetPublicKeyBytes(privKey *ecdsa.PrivateKey) []byte {
	pubKey := privKey.PublicKey
	pubKeyBytes := append(pubKey.X.Bytes(), pubKey.Y.Bytes()...)
	return pubKeyBytes
}

// TODO check if kp.SigningKey is nil
func (kp *KeyPair) PrintPublicKey() error {
	if kp.SigningKey == nil {
		return ErrKeyNotFound
	}
	fmt.Println("your public key is: ", hex.EncodeToString(kp.GetPublicKeyBytes(kp.SigningKey)))
	return nil
}

func (kp *KeyPair) StorePrivateKey(path string) error {
	pubKeyBytes := kp.GetPublicKeyBytes(kp.SigningKey)
	err := os.WriteFile(path, pubKeyBytes, 0644)
	if err != nil {
		return err
	}

	fmt.Println("your private key was stored at ", path)

	return nil
}

func ViewPrivateKey(path string, pwd string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	privKey := hex.EncodeToString(data)
	fmt.Println("your private key is: ", privKey)
	return nil
}
