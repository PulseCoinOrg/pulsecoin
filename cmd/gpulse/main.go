package main

import (
	"github.com/PulseCoinOrg/pulsecoin/console"
	"github.com/charmbracelet/log"
)

func main() {
	console := console.New()
	if err := console.Run(); err != nil {
		log.Error(err.Error())
	}
}
