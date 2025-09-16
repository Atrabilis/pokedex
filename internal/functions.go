package internal

import (
	"fmt"
	"os"
	"strings"
)

func CleanInput(input string) []string {
	input = strings.ToLower(input)
	return strings.Fields(input)
}
func CommandHelp(string) error {
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range Commands {
		fmt.Println(command.Name+":", command.Description)
	}
	return nil
}
func CommandExit(string) error {
	os.Exit(0)
	return nil
}
