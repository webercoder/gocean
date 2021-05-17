package command

import "flag"

// Handler is a general command-line handler interface used by this application.
type Handler interface {
	GetFlagSet(command string) (*flag.FlagSet, error)
	HandleCommand(command string) error
	Usage(err ...error)
}
