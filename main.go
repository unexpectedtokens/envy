package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

const CREATE_NEW_ENV = "Create a new env file"

func createNewENVFile() {
	fmt.Println("Creating new env file")
	prompt := promptui.Prompt{
		Label: "Type in the name the file should have. Remember that the name should start with a . and should end with .env if you wish to have the ability to update the file later on with Envy",
	}
	prompt.Run()

}

func editExistingENVFile() {

}

func main() {
	fmt.Println("Scanning for existing .env files...")
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dirEntries, err := os.ReadDir(wd)
	if err != nil {
		panic(err)
	}
	existingEnv := []string{}
	for _, de := range dirEntries {
		fileName := de.Name()
		if string(fileName[0]) == "." {
			if strings.Contains(de.Name(), "env") {
				existingEnv = append(existingEnv, fileName)

			}
		}
	}
	if len(existingEnv) > 0 {
		fmt.Println("Existing env file(s) found. Pick one of the files to edit or create a new one:")
		prompt := promptui.Select{
			Label: "Select one of the options by going up or down with the arrow keys and confirming with the enter key",
			Items: append(existingEnv, CREATE_NEW_ENV),
		}

		number, value, err := prompt.Run()

		if err != nil {
			panic(err)
		}
		if value == CREATE_NEW_ENV {
			createNewENVFile()
		}
		fmt.Println(number)
	} else {
		fmt.Println("No existing env files found")
		createNewENVFile()
	}

}
