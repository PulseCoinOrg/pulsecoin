package accounts

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
)

type KeyPair struct {
	PublicKey  *ecdsa.PublicKey
	PrivateKey *ecdsa.PrivateKey
}

func GenerateKeys() (*KeyPair, error) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	pubKey := privKey.PublicKey
	return &KeyPair{
		PublicKey:  &pubKey,
		PrivateKey: privKey,
	}, nil
}

func printWarning() {
	fmt.Println("\n")
	fmt.Println("\t\tYou NEVER share your private key with anyone")
	fmt.Println("\t\tYour public key is sharable, and this is how people send you funds")
	fmt.Println("\n")
}

func (kp *KeyPair) PrintPublicKey() {
	pubKeyBytes := append(kp.PublicKey.X.Bytes(), kp.PublicKey.Y.Bytes()...)
	pubKeyString := hex.EncodeToString(pubKeyBytes)
	fmt.Printf("Your public key is > %s\n", pubKeyString)
}

func (kp *KeyPair) StorePrivateKey(path string) error {
	privKeyBytes := kp.PrivateKey.D.Bytes()
	err := os.WriteFile(path, privKeyBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

// TODO add password support
func ViewPrivateKey(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	printWarning()
	fmt.Printf("> %s\n", hex.EncodeToString(data))

	return nil
}
