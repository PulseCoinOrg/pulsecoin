package console

import (
	"fmt"
	"os"

	"github.com/PulseCoinOrg/pulsecoin/accounts"
	"github.com/PulseCoinOrg/pulsecoin/common"
)

func Help() {
	fmt.Println("\t\tgpulse> [command] [...args]")
}

func Exit() {
	os.Exit(1)
}

func GenKeyPair(pwd string) {
	keys, err := accounts.NewKeyPair(pwd)
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

func ViewKey(path string, pwd string) {
	err := accounts.ViewPrivateKey(path, pwd)
	if err != nil {
		panic(err)
	}
}
