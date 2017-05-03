package main

import (
	"github.com/WindomZ/gituser/gituser"
	"github.com/WindomZ/go-commander"
)

func main() {
	// gituser
	commander.Program.
		Version("v1.2.0")

	// gituser add [options] <user> <name> <email>
	commander.Program.
		Command("add <user> <name> <email>").
		Description("add user configuration").
		Option("--private-github", "private email address for GitHub").
		Action(gituser.AddAction)

	// gituser remove <user>
	commander.Program.
		Command("remove <user>").
		Aliases([]string{"rm"}).
		Description("remove user configuration").
		Action(gituser.RemoveAction)

	// gituser list
	commander.Program.
		Command("list").
		Aliases([]string{"ls"}).
		Description("list user configuration").
		Action(gituser.ListAction)

	// gituser set [options] <user>
	commander.Program.
		Command("set <user>").
		Description("set local git user configuration from <user>").
		Option("--private-github", "private email address for GitHub").
		Action(gituser.SetAction)

	// gituser unset [options]
	commander.Program.
		Command("unset").
		Description("unset local git user configuration").
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
