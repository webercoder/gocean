package lib

import "flag"

// GoceanCommandHandler .
type CommandHandler interface {
	GetFlagSet(command string) (*flag.FlagSet, error)
	HandleCommand(command string) error
	Usage(err ...error)
}
