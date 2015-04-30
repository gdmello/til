package command

type InitCommand struct {
}

func (c *InitCommand) Run(args []string) int {
	return 0
}

func (c *InitCommand) Help() string {
	return "Initialize the TIL repository, default to the current directory."
}

func (c *InitCommand) Synopsis() string {
	return "Initialize the TIL repository"
}
