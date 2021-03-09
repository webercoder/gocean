package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/webercoder/gocean/noaa/stations"
	"github.com/webercoder/gocean/noaa/tidesandcurrents"
)

func usage(msg string) {
	if len(msg) > 0 {
		fmt.Printf("%s\n", msg)
	}
	fmt.Printf("Usage:\n\t%s stations postcode\n\t%s tidesandcurrents station-id\n", os.Args[0], os.Args[0])
	os.Exit(1)
}

type commandHandlerContainer struct {
	flagSet *flag.FlagSet
	handler GoceanCommandHandler
}

func main() {
	handlers := map[string]commandHandlerContainer{
		"stations": {
			flagSet: flag.NewFlagSet("stations", flag.ExitOnError),
			handler: &stations.CommandHandler{},
		},
		"tidesandcurrents": {
			flagSet: flag.NewFlagSet("tidesandcurrents", flag.ExitOnError),
			handler: &tidesandcurrents.CommandHandler{},
		},
	}

	if len(os.Args) < 2 {
		usage("")
	}

	handlerContainer, ok := handlers[os.Args[1]]
	if !ok {
		usage(fmt.Sprintf("%s is not a valid command", os.Args[1]))
	}

	if err := handlerContainer.flagSet.Parse(os.Args[2:]); err != nil {
		usage("Unable to parse command line options")
	}

	err := handlerContainer.handler.HandleCommand(handlerContainer.flagSet.Arg(0))
	if err != nil {
		usage(fmt.Sprint(err))
	}
}
