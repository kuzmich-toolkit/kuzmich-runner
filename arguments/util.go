package arguments

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

func parseBooleanFlag(command *cobra.Command, name string) (bool, error) {
	return strconv.ParseBool(command.Flag(name).Value.String())
}

func parseMapFlag(command *cobra.Command, name string) map[string]string {
	value := command.Flag(name).Value.String()
	arrayValue := strings.Split(value, " ")

	if len(arrayValue)%2 != 0 {
		log.Fatalf(
			"Error during parsing map falg with name: %s. You should pass even number of words",
			name,
		)
	}

	result := map[string]string{}

	for i := 0; i < len(arrayValue); i += 2 {
		result[arrayValue[i]] = arrayValue[i+1]
	}

	return result
}
