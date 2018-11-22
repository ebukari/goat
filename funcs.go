package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Stub creation related functions and variables

var header = "package %s\n\nimport \"fmt\"\n\n"
var funcStub = "func %s() {\n    fmt.Println(\"Hello World!\")\n}"

func createStub(path, name string, stubtype bool) error {
	stubFile, err := createStubFile(path, name, stubtype)
	defer stubFile.Close()
	if err != nil {
		return err
	}
	fmt.Fprintf(stubFile, formatStubCode(name, stubtype))
	return err
}
func createStubFile(path, name string, pkg bool) (*os.File, error) {
	stubName := "main.go"
	if pkg {
		stubName = name + ".go"
	}
	stubPath := filepath.Join(path, stubName)
	return os.Create(stubPath)
}
func formatStubCode(name string, isPkg bool) string {
	funcCode := fmt.Sprintf(funcStub, name)
	if !isPkg {
		funcCode = fmt.Sprintf(funcStub, "main")
	}
	return fmt.Sprintf(header, name) + funcCode
}

// File path related functions
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func userInput(prompt string) string {
	var uInput string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&uInput)
	fatalCheckMe("reading user input", err)
	return uInput
}

// askYesNo question prompts the user and returns true if the user responded
// with "y" and false if the user responded with "n"
// Any other value cause the question to be asked again.
func askYesNo(prompt string) bool {
	uInput := strings.TrimSpace(userInput(prompt))
	if uInput == "y" {
		return true
	} else if uInput == "n" {
		return false
	}
	Logger.Println("Wrong input. Expecting either 'y' or 'n'.")
	return askYesNo(prompt)
}

// Error Handling Functions

// fatalCheck alerts that there was an error without
// giving any extra information and the exits.
func fatalCheck(err error) {
	if err != nil {
		Logger.Fatalln("An error occured. Exiting...")
	}
}

// fatalCheckm handles errors by showing a message and obscuring the actual error
func fatalCheckm(message string, err error) {
	if err != nil {
		Logger.Fatalf("error occured while %s\n", message)
	}
}

// fatalCheckMe handles error by adding an extra message to explain the error
func fatalCheckMe(message string, err error) {
	if err != nil {
		Logger.Fatalf("error occured while %s: %v\n", message, err)
	}
}
