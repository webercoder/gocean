package main

// GoceanCommandHandler .
type GoceanCommandHandler interface {
	HandleCommand(arg string) error
}
