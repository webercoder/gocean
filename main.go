package main

import (
	"errors"
	"fmt"
	"os"

	tidesandcurrents "github.com/webercoder/gocean/coops"
	"github.com/webercoder/gocean/lib"
	"github.com/webercoder/gocean/stations"
)

func usage(handlers map[string]lib.CommandHandler, msg ...string) {
	if len(msg) > 0 {
		for i := 0; i < len(msg); i++ {
			fmt.Println(msg)
		}
	}

	fmt.Println("Usage:")
	for _, handler := range handlers {
		handler.Usage()
	}

	os.Exit(1)
}

func main() {
	handlers := map[string]lib.CommandHandler{
		"stations":         stations.NewCommandHandler(),
		"tidesandcurrents": tidesandcurrents.NewTidesAndCurrentsCommandHandler(),
	}

	if len(os.Args) < 2 {
		usage(handlers)
	}

	command := os.Args[1]
	handler, ok := handlers[command]
	if !ok {
		usage(handlers, fmt.Sprintf("Command %s is not a valid top-level command", command))
	}

	var subcommand string
	if len(os.Args) > 2 {
		subcommand = os.Args[2]
	}

	fset, err := handler.GetFlagSet(subcommand)
	if err != nil {
		handler.Usage(err)
		os.Exit(1)
	}

	if err := fset.Parse(os.Args[1:]); err != nil {
		handler.Usage(errors.New("Unable to parse command-line options"))
		os.Exit(1)
	}

	err = handler.HandleCommand(subcommand)
	if err != nil {
		fmt.Println(err)
	}
}
