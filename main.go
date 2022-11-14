package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
)

var workingDirectory string

func main() {
	fmt.Println("Scanning for existing .env files...")
	var err error
	workingDirectory, err = os.Getwd()

	if err != nil {
		panic(err)
	}

	dirEntries, err := os.ReadDir(workingDirectory)
	if err != nil {
		panic(err)
	}
	existingEnv := []string{}
	fmt.Println("Found the following env files. Pick the one you wish to load into this environment:")

	for _, de := range dirEntries {
		fileName := de.Name()
		if string(fileName[0]) == "." {
			if strings.Contains(de.Name(), "env") {
				existingEnv = append(existingEnv, fileName)
				// fmt.Printf("- %s\n", fileName)
			}
		}
	}
	fmt.Println()
	prompt := promptui.Select{
		Label: "Select one of the options by going up or down with the arrow keys and confirming with the enter key",
		Items: existingEnv,
	}

	_, value, err := prompt.Run()

	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	godotenv.Load(value)

}
