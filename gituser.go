package main

import (
	"github.com/WindomZ/gituser/gituser"
	"github.com/WindomZ/go-commander"
)

func main() {
	// gituser
	commander.Program.
		Version("v0.0.0")

	// gituser add [options] <user> <name> <email>
	commander.Program.
		Command("add <user> <name> <email>").
		Description("add user configuration").
		Option("--private-github", "private email address for GitHub").
		Action(gituser.AddAction)

	// gituser list
	commander.Program.
		Command("list").
		Description("list user configuration").
		Action(gituser.ListAction)

	// gituser set [options] <user>
	commander.Program.
		Command("set <user>").
		Description("set local git user configuration from <user>").
		Option("-g --global", "set global git user configuration").
		Option("--private-github", "private email address for GitHub").
		Action(gituser.SetAction)

	// gituser unset [options]
	commander.Program.
		Command("unset").
		Description("unset local git user configuration").
		Option("-g --global", "unset global git user configuration").
		Action(gituser.UnsetAction)

	commander.Program.Annotation(
		"Argument",
		[]string{
			commander.Format.Description("<user>",
				"the name of the configure user information"),
			commander.Format.Description("<name>",
				"the name of the git username"),
			commander.Format.Description("<email>",
				"the address of the git email"),
		},
	)

	if _, err := commander.Program.Parse(); err != nil {
		panic(err)
	}
}
