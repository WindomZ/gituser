package gituser

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

var RemoveAction = func(c commander.Context) error {
	users, err := readConfig()
	if err != nil {
		return err
	}
	strUser := c.MustString("<user>")
	if len(strUser) == 0 || !users.Remove(strUser) {
		fmt.Println("Not found:", strUser)
		return nil
	}
	if err := writeConfigs(users); err != nil {
		return err
	}
	fmt.Println("Success!")
	return nil
}
