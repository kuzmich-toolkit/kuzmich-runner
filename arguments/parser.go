package arguments

import (
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"
	"log"
)

func parseSharedFlags(command *cobra.Command, arguments *CommonRunnerArguments) error {
	var err error

	arguments.Apk = command.Flag("apk").Value.String()
	arguments.TestApk = command.Flag("test-apk").Value.String()
	arguments.Package = command.Flag("package").Value.String()
	arguments.TestPackage = command.Flag("test-package").Value.String()
	arguments.Output = command.Flag("output").Value.String()

	verbose, verboseErr := parseBooleanFlag(command, "verbose")
	if verboseErr != nil {
		multierror.Append(err, verboseErr)
	} else {
		arguments.Verbose = verbose
	}

	return err
}

func ParseCommonRunFlags(command *cobra.Command) *CommonRunnerArguments {
	result := NewCommonRunnerArguments()

	err := parseSharedFlags(command, result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
