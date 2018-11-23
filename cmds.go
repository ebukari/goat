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
	Flags:  []cli.Flag{isPkgFlag, workDirFlag, editorFlag},
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

// workDirFlag allows user to specify the `WORKDIR` for the project
var workDirFlag = cli.StringFlag{
	Name:        "w",
	Value:       ".",
	Usage:       "creat project in `WORKDIR`",
	Destination: &workDir,
}

// editorFlag is for specifying the `EDITOR` to open the project with after
// creating it
var editorFlag = cli.StringFlag{
	Name:        "editor, e",
	Usage:       "use `PATH_TO_EDITOR` use when opening project",
	Destination: &editorPath,
}

// isPkgFlag specifies whether to set up the project as a package
var isPkgFlag = cli.BoolFlag{
	Name:        "package, p",
	Usage:       "if supplied, sets up the project as a package",
	Destination: &isPkg,
}
