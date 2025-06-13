package prompt

import (
	"fmt"
	"reflect"
)

type Command struct {
	Name string
	Func interface{}
}

type Dispatcher struct {
	commands map[string]*Command
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{commands: make(map[string]*Command)}
}

func (d *Dispatcher) Register(cmd *Command) {
	d.commands[cmd.Name] = cmd
}

func (d *Dispatcher) Call(name string, args ...string) error {
	cmd, exists := d.commands[name]
	if !exists {
		return fmt.Errorf("unknown command: %s", name)
	}

	fn := reflect.ValueOf(cmd.Func)
	if fn.Kind() != reflect.Func {
		return fmt.Errorf("not a function")
	}

	if fn.Type().NumIn() != len(args) {
		return fmt.Errorf("expected %d args, got %d", fn.Type().NumIn(), len(args))
	}

	// Convert args to reflect.Value
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	fn.Call(in)
	return nil
}
