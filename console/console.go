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
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "keygen", Func: GenKeyPair})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "viewprivkey", Func: ViewSigningKey})
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: "exit", Func: Exit})
	err := c.TermPrompt.PromptInput("gpulse> ")
	if err != nil {
		return err
	}
	return nil
}
