package gituser

import "github.com/WindomZ/go-commander"

var SetAction = func(c commander.Context) error {
	println("set...", c.MustString("<user>"))
	return nil
}
