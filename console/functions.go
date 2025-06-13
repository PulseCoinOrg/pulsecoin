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
	kp, err := accounts.NewKeyPair(pwd)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = kp.PrintPublicKey()
	if err != nil {
		fmt.Println(err.Error())
	}

	keyHash := common.Sha256Hash([]byte("private-key-hash-file-name")[:10]) // TODO make a random string generator for this
	err = kp.StorePrivateKey(fmt.Sprintf("gpulse-secret-key_%s", keyHash.String()))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ViewSigningKey(path string, pwd string) {
	err := accounts.ViewPrivateKey(path, pwd)
	if err != nil {
		fmt.Println(err.Error())
	}
}
