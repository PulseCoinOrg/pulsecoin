package main

import "github.com/PulseCoinOrg/pulsecoin/console"

func main() {
	cons := console.New()
	if err := cons.Run(); err != nil {
		panic(err)
	}
}
