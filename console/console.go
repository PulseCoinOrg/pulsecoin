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
	c.TermPrompt.Dispatcher.Register(&prompt.Command{Name: commands[0], Func: GenKeyPair})
	err := c.TermPrompt.PromptInput("gpulse> ")
	if err != nil {
		return err
	}
	return nil
}
