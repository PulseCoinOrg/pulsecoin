package console

import (
	"fmt"

	"github.com/PulseCoinOrg/pulsecoin/accounts"
	"github.com/PulseCoinOrg/pulsecoin/common"
)

func GenKeyPair() {
	kp := &accounts.KeyPair{}

	keys, err := kp.New()
	if err != nil {
		panic(err)
	}

	err = keys.PrintPublicKey()
	if err != nil {
		panic(err)
	}

	keyHash := common.Sha256Hash([]byte("private-key-hash-file-name"))
	err = keys.StorePrivateKey(fmt.Sprintf("gpulse-secret-key_%s", keyHash.String()))
	if err != nil {
		panic(err)
	}
}
