package tidesandcurrents

import (
	"fmt"

	"github.com/webercoder/gocean/noaa/tidesandcurrents/predictions"
	"github.com/webercoder/gocean/noaa/tidesandcurrents/utils"
)

// CommandHandler .
type CommandHandler struct{}

// GetSupportedCommands .
func (pch *CommandHandler) GetSupportedCommands() []string {
	return []string{"predictions"}
}

// HandleCommand .
func (pch *CommandHandler) HandleCommand(arg string) error {
	tidesClient := utils.NewClient()
	results, err := predictions.Retrieve(tidesClient, arg, 24)
	if err != nil {
		return fmt.Errorf("Could not load predictions for station")
	}
	predictions.PrintTabDelimited(arg, results)
	return nil
}
