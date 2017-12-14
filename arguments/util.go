package arguments

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func parseBooleanFlag(command *cobra.Command, name string) bool {
	value, err := strconv.ParseBool(command.Flag(name).Value.String())
	if err != nil {
		log.Fatalf("Error during parsing boolean flag with name: %s", name)
	}

	return value
}
