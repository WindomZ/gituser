package gituser

import "github.com/WindomZ/go-commander"

var AddAction = func(c commander.Context) error {
	println("add...", c.MustString("<user>"), c.MustString("<name>"), c.MustString("<email>"))
	return nil
}
