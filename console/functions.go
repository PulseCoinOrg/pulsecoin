package console

import (
	"fmt"
	"os"

	"github.com/PulseCoinOrg/pulsecoin/accounts"
	"github.com/charmbracelet/log"
)

func Help() {
	fmt.Println("\t\tgpulse> [command] [...args]")
}

func Exit() {
	os.Exit(1)
}

func WalletNew(path string) {
	_, err := accounts.New(path)
	if err != nil {
		log.Error(err.Error())
	}
}

func PrivKeyView(path string) {
	err := accounts.ViewPrivateKey(path)
	if err != nil {
		log.Error(err.Error())
	}
}
