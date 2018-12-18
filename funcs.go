package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// File modes
const (
	OS_READ        = 04
	OS_WRITE       = 02
	OS_EX          = 01
	OS_USER_SHIFT  = 6
	OS_GROUP_SHIFT = 3
	OS_OTH_SHIFT   = 0

	OS_USER_R   = OS_READ << OS_USER_SHIFT
	OS_USER_W   = OS_WRITE << OS_USER_SHIFT
	OS_USER_X   = OS_EX << OS_USER_SHIFT
	OS_USER_RW  = OS_USER_R | OS_USER_W
	OS_USER_RWX = OS_USER_RW | OS_USER_X

	OS_GROUP_R   = OS_READ << OS_GROUP_SHIFT
	OS_GROUP_W   = OS_WRITE << OS_GROUP_SHIFT
	OS_GROUP_X   = OS_EX << OS_GROUP_SHIFT
	OS_GROUP_RW  = OS_GROUP_R | OS_GROUP_W
	OS_GROUP_RWX = OS_GROUP_RW | OS_GROUP_X

	OS_OTH_R   = OS_READ << OS_OTH_SHIFT
	OS_OTH_W   = OS_WRITE << OS_OTH_SHIFT
	OS_OTH_X   = OS_EX << OS_OTH_SHIFT
	OS_OTH_RW  = OS_OTH_R | OS_OTH_W
	OS_OTH_RWX = OS_OTH_RW | OS_OTH_X

	OS_ALL_R   = OS_USER_R | OS_GROUP_R | OS_OTH_R
	OS_ALL_W   = OS_USER_W | OS_GROUP_W | OS_OTH_W
	OS_ALL_X   = OS_USER_X | OS_GROUP_X | OS_OTH_X
	OS_ALL_RW  = OS_ALL_R | OS_ALL_W
	OS_ALL_RWX = OS_ALL_RW | OS_GROUP_X
)

// Stub creation related functions and variables

var header = "package %s\n\nimport \"fmt\"\n\n"
var funcStub = "func %s() {\n    fmt.Println(\"Hello World!\")\n}"

// create stub for project and add the default placeholder code to it
func createStub(path, name string, stubtype bool) error {
	stubFile, err := createStubFile(path, name, stubtype)
	defer stubFile.Close()
	if err != nil {
		return err
	}
	fmt.Fprintf(stubFile, formatStubCode(name, stubtype))
	return err
}

// createStubFile creates stub file and returns a the handle to the file and
// any error that occured while creating it
func createStubFile(path, name string, pkg bool) (*os.File, error) {
	stubName := "main.go"
	if pkg {
		stubName = name + ".go"
	}
	stubPath := filepath.Join(path, stubName)
	return os.Create(stubPath)
	// return os.OpenFile(stubPath, os.O_RDWR|os.O_EXCL|os.O_CREATE|os.O_APPEND, os.ModeSetuid|os.ModeSetgid|os.ModeExclusive)
}

// formatStubCode formats the default code for stub files and return it as a string
func formatStubCode(name string, isPkg bool) string {
	funcCode := fmt.Sprintf(funcStub, name)
	if !isPkg {
		funcCode = fmt.Sprintf(funcStub, "main")
	}
	return fmt.Sprintf(header, name) + funcCode
}

// File path related functions

// pathExists returns true if a path already exists
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// userInput prompts the user for entry of information and returns the information
// as a string
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

// fatalCheckm obscures the actual error with another message
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
