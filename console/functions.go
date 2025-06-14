package console

import (
	"fmt"
	"os"
	"os/exec"
	"time"

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

func PrivKeyRevoke(path string) {
	var confirmed string
	fmt.Println("Are you sure you would like to revoke your private key? [y/n]: ")
	fmt.Scan(&confirmed)

	if confirmed == "y" {
		fmt.Println("revoking private key...")
		time.Sleep(3 * time.Second)
		cmd := exec.Command("sudo", "rm", path)
		if err := cmd.Run(); err != nil {
			log.Error(err.Error())
		}
		fmt.Println(cmd)
		time.Sleep(2 * time.Second)
		fmt.Println("private key has been revoked and deleted from >", path)
	} else if confirmed == "n" {
		os.Exit(1)
	}
}
