package main

import (
	"fmt"
	"os"

	"github.com/webercoder/gocean/command"
)

func usage(handlers map[string]command.Handler, err ...string) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
	}

	for key := range handlers {
		fmt.Printf("  %v\n", key)
	}

	os.Exit(1)
}

func main() {
	handlers := map[string]command.Handler{
		"stations": command.NewStationsCommandHandler(),
		"coops":    command.NewCOOPSCommandHandler(),
	}

	if len(os.Args) < 2 {
		usage(handlers, "Please provide a subcommand")
	}

	command := os.Args[1]
	handler, ok := handlers[command]
	if !ok {
		usage(handlers, fmt.Sprintf("Command %s is not a valid top-level command", command))
	}

	// var subcommand string
	// if len(os.Args) > 2 {
	// 	subcommand = os.Args[2]
	// }

	// fset, err := handler.GetFlagSet(subcommand)
	// if err != nil {
	// 	handler.Usage(err)
	// }

	// if err := fset.Parse(os.Args[1:]); err != nil {
	// 	handler.Usage(errors.New("unable to parse command-line options"))
	// }

	err := handler.HandleCommand()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
