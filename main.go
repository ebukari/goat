package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

// Logger is used to output information to the command line
var Logger = log.New(os.Stdout, app.Name+": ", 0)

func init() {
	app.Name = "goat"
	app.Author = "Ezekiel N. Bukari"
	app.Email = "enbukari@gmail.com"
	app.Version = "0.0.1"
	app.Usage = "stub new Go projects"
	app.UsageText = `GOlang App Templater[GOAT] is a tool to create and manage Go projects`
}

func main() {
	app.Commands = []cli.Command{
		newCmd,
		configCmd,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// CreateProject initialises a project default flags or flags passed to the command
func CreateProject(c *cli.Context) error {
	if c.NArg() == 0 {
		Logger.Fatal("Please supply a name for the project")
	}
	projectName := strings.ToLower(c.Args().Get(0))
	projectPath, err := filepath.Abs(workDir)
	fatalCheckMe("resolving home directory of project", err)
	projectHome := filepath.Join(projectPath, projectName)
	if pathExists(projectHome) {
		if !askYesNo("Directory exists. Overwrite? (y/n): ") {
			Logger.Fatal("Exiting...")
		}
		err = os.RemoveAll(projectHome)
		fatalCheckMe("removing existing directory", err)
	}
	Logger.Printf("Creating project in %s", projectHome)
	err = os.MkdirAll(projectHome, os.ModeDir)
	err = createStub(projectHome, projectName, isPkg)
	if editorPath != "" {
		cmd := exec.Command(editorPath, projectHome)
		_ = cmd.Start()
	}
	return err
}
