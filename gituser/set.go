package gituser

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

var SetAction = func(c commander.Context) error {
	setDebugMode(c.MustBool("--debug"))

	users, err := readConfig()
	if err != nil {
		return err
	}
	strUser := c.MustString("<user>")
	user, ok := users.Find(strUser)
	if !ok || user == nil {
		fmt.Println("Not found:", strUser)
		return nil
	}
	if _, err := commander.Exec.ExecStdCommand(
		"git", "config", "user.name", user.Name); err != nil {
		return err
	}
	if _, err := commander.Exec.ExecStdCommand(
		"git", "config", "user.email", user.Email); err != nil {
		return err
	}
	fmt.Println("Success!")
	return nil
}
