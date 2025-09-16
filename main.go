package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atrabilis/pokedex/internal"
)

func main() {
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := internal.CleanInput(scanner.Text())
		for _, command := range internal.Commands {
			if input[0] == command.Name {
				command.FallbackFuction(input[0])
			}
		}
	}
}
