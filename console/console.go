package console

import (
	"fmt"

	"github.com/PulseCoinOrg/pulsecoin/console/prompt"
	"github.com/PulseCoinOrg/pulsecoin/core"
)

var commands = []string{
	"keygen",
}

type Console struct {
	TermPrompt *prompt.TerminalPrompt
	chain      *core.BlockChain
}

func New(chain *core.BlockChain) *Console {
	dispatch := prompt.NewDispatcher()
	return &Console{
		TermPrompt: &prompt.TerminalPrompt{Dispatcher: dispatch},
		chain:      chain,
	}
}

func (c *Console) Run() error {
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "help", Func: Help})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "exit", Func: Exit})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "wallet-new", Func: WalletNew})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "privkey-view", Func: PrivKeyView})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "privkey-revoke", Func: PrivKeyRevoke})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "block-search", Func: func(hash string) {
		BlockSearch(c.chain, hash)
	}})

	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "cmd-dump", Func: func() {
		c.TermPrompt.Dispatcher.DumpCommands()
	}})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "cls", Func: func() {
		fmt.Print("\033[H\033[2J")
	}})

	err := c.TermPrompt.PromptInput("gpulse> ")
	if err != nil {
		return err
	}
	return nil
}
