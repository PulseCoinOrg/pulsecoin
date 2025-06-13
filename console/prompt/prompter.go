package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Prompter interface {
	PromptInput(prompt string) (string, error)
}

type TerminalPrompt struct {
	Dispatcher *Dispatcher
	commands   []*Command
}

func (tp *TerminalPrompt) PromptInput(prompt string) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		cmdName := tokens[0]
		args := tokens[1:]

		err = tp.Dispatcher.Call(cmdName, args...)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
	}
}
