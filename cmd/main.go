package main

import "github.com/PulseCoinOrg/pulsecoin/console"

func main() {
	console := console.New()
	if err := console.Run(); err != nil {
		panic(err)
	}
}
