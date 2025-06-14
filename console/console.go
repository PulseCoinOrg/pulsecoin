package console

import "github.com/PulseCoinOrg/pulsecoin/console/prompt"

var commands = []string{
	"keygen",
}

type Console struct {
	TermPrompt *prompt.TerminalPrompt
}

func New() *Console {
	dispatch := prompt.NewDispatcher()
	return &Console{
		TermPrompt: &prompt.TerminalPrompt{Dispatcher: dispatch},
	}
}

func (c *Console) Run() error {
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "help", Func: Help})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "exit", Func: Exit})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "wallet-new", Func: WalletNew})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "privkey-view", Func: PrivKeyView})
	err := c.TermPrompt.PromptInput("gpulse> ")
	if err != nil {
		return err
	}
	return nil
}
