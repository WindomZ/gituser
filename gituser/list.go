package gituser

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

var ListAction = func(c commander.Context) error {
	setDebugMode(c.MustBool("--debug"))

	users, err := readConfig()
	if err != nil {
		return err
	}
	for _, user := range users.Users {
		fmt.Println(fmt.Sprintf("%s - %s<%s>", user.User, user.Name, user.Email))
	}
	return nil
}
