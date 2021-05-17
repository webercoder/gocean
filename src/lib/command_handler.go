package lib

import "flag"

// CommandHandler is a general command-line handler interface used by this application.
type CommandHandler interface {
	GetFlagSet(command string) (*flag.FlagSet, error)
	HandleCommand(command string) error
	Usage(err ...error)
}
