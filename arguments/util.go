package arguments

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func parseBooleanFlag(command *cobra.Command, name string) (bool, error) {
	return strconv.ParseBool(command.Flag(name).Value.String())
}

func parseMapFlag(command *cobra.Command, name string) (map[string]string, error) {
	result := map[string]string{}

	value := command.Flag(name).Value.String()

	if len(value) == 0 {
		return result, nil
	}

	arrayValue := strings.Split(value, " ")
	if len(arrayValue)%2 != 0 {
		return nil, errors.New(
			fmt.Sprintf(
				"Error during parsing map flag with name: %s. You should pass even number of words like this: "+
					"--%s \"argument1 value1 argument2 value2\"",
				name,
				name,
			),
		)
	}

	for i := 0; i < len(arrayValue); i += 2 {
		result[arrayValue[i]] = arrayValue[i+1]
	}

	return result, nil
}
