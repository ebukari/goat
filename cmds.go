package main

import "github.com/urfave/cli"

// placeHolder is for default action for new commands
func placeHolder(c *cli.Context) error {
	return nil
}

// Commands
var newCmd = cli.Command{
	Name:   "new",
	Usage:  "create a new project",
	Action: CreateProject,
	Flags: []cli.Flag{
		isPkgFlag,
		workDirFlag,
		editorFlag,
	},
}
var configCmd = cli.Command{
	Name:   "config",
	Usage:  "configure default options for creating projects",
	Action: placeHolder,
}

// Flag variables
var workDir string

// path of the editor to use
var editorPath string

// setup the project as a package
var isPkg bool

// Flag Declarations
var workDirFlag = cli.StringFlag{
	Name:        "w",
	Value:       ".",
	Usage:       "creat project in `WORKDIR`",
	Destination: &workDir,
}
var editorFlag = cli.StringFlag{
	Name:        "editor, e",
	Usage:       "use `PATH_TO_EDITOR` use when opening project",
	Destination: &editorPath,
}
var isPkgFlag = cli.BoolFlag{
	Name:        "package, p",
	Usage:       "if supplied, sets up the project as a package",
	Destination: &isPkg,
}
