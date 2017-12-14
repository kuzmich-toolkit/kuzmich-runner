package arguments

import (
	"github.com/spf13/cobra"
)

func parseSharedFlags(command *cobra.Command, arguments *CommonRunnerArguments) {
	arguments.Apk = command.Flag("apk").Value.String()
	arguments.TestApk = command.Flag("test-apk").Value.String()
	arguments.Package = command.Flag("package").Value.String()
	arguments.TestPackage = command.Flag("test-package").Value.String()
	arguments.Output = command.Flag("output").Value.String()

	arguments.Verbose = parseBooleanFlag(command, "verbose")

}

func ParseCommonRunFlags(command *cobra.Command) *CommonRunnerArguments {
	result := NewCommonRunnerArguments()

	parseSharedFlags(command, result)

	return result
}
