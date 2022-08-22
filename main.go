package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

const CREATE_NEW_ONE = "Create a new env file"

func main() {
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
			Items: append(existingEnv, CREATE_NEW_ONE),
		}
		prompt.Run()
	}

}
