package coops

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/webercoder/gocean/src/coops"
)

// BaseCommandHandler handles predictions commands.
type BaseCommandHandler struct {
	clientConfig *ClientConfig
	flagSet      *flag.FlagSet
	product      coops.Product
}

// NewBaseCommandHandler creates a new Tides and Currents CommandHandler.
func NewBaseCommandHandler(prod coops.Product) *BaseCommandHandler {
	clientConfig := NewClientConfig()
	return &BaseCommandHandler{
		clientConfig: clientConfig,
		flagSet:      clientConfig.GetFlagSet(prod.String(), flag.ExitOnError),
		product:      prod,
	}
}

// GetRequestOptions provides coops.ClientRequestOptions for the CO-OPS API request.
func (b *BaseCommandHandler) GetRequestOptions() []coops.ClientRequestOption {
	reqOptions, err := b.clientConfig.ToRequestOptions()
	if err != nil {
		b.Usage(err)
	}

	return append(reqOptions, coops.WithProduct(b.product))
}

// HandleCommand must be overridden.
func (b *BaseCommandHandler) HandleCommand() {
	panic(errors.New("please override the base command handler's HandleCommand method"))
}

// ParseFlags parses command-line arguments and exits if they are invalid.
func (b *BaseCommandHandler) ParseFlags() {
	if err := b.flagSet.Parse(os.Args[3:]); err != nil {
		b.Usage(errors.New("unable to parse command-line options"))
	}
}

// Usage prints how to use this command.
func (b *BaseCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
	}

	b.flagSet.Usage()
	os.Exit(1)
}
