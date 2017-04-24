package gituser

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

var UnsetAction = func(c commander.Context) error {
	if _, err := commander.Exec.ExecStdCommand(
		"git", "config", "--unset", "user.name"); err != nil {
		return err
	}
	if _, err := commander.Exec.ExecStdCommand(
		"git", "config", "--unset", "user.email"); err != nil {
		return err
	}
	fmt.Println("Success!")
	return nil
}
