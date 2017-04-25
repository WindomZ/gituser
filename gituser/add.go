package gituser

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

var AddAction = func(c commander.Context) error {
	user := &_GitUser{
		User:  c.MustString("<user>"),
		Name:  c.MustString("<name>"),
		Email: c.MustString("<email>"),
	}
	if c.MustBool("--private-github") {
		user.Email = fmt.Sprintf("%s@users.noreply.github.com", user.Name)
	}
	if err := user.Valid(); err != nil {
		return err
	} else if err := writeConfig(user); err != nil {
		return err
	}
	fmt.Println("Success!")
	return nil
}
