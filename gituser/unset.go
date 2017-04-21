package gituser

import "github.com/WindomZ/go-commander"

var UnsetAction = func(c commander.Context) error {
	println("unset...")
	return nil
}
