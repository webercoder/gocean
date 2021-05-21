package command

// Handler is a general command-line handler interface used by this application.
type Handler interface {
	HandleCommand() error
	Usage(err ...error)
}
