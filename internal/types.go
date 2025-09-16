package internal

type CommandLine struct {
	Name            string
	Description     string
	FallbackFuction func(string) error
}

var Commands = map[string]CommandLine{}

func init() {
	Commands = map[string]CommandLine{
		"help": {
			Name:            "help",
			Description:     "Displays a help message",
			FallbackFuction: CommandHelp,
		},
		"exit": {
			Name:            "exit",
			Description:     "Exit the Pokedex",
			FallbackFuction: CommandExit,
		},
	}
}
