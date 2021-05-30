package command

import "github.com/webercoder/gocean/src/coops"

// Handler is a general command-line handler interface used by this application.
type Handler interface {
	HandleCommand() error
	ParseFlags()
	Usage(err ...error)
}

// CoopsHandler embeds Handler but adds a COOPS request option method for providing ClientRequestOptions.
type CoopsHandler interface {
	Handler
	GetRequestOptions() []coops.ClientRequestOption
}
