package gituser

import "github.com/WindomZ/go-commander"

var UnsetArgvAction = func(c commander.Context) error {
	println("unset...")
	return nil
}
