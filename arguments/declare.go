package arguments

import "github.com/spf13/cobra"

func declareSharedFlags(command *cobra.Command) {
	command.Flags().String("apk", "", "Android application path")
	command.MarkFlagRequired("apk")

	command.Flags().String("test-apk", "", "Android test application path")
	command.MarkFlagRequired("test-apk")

	command.Flags().String("package", "", "Target application package")
	command.MarkFlagRequired("package")

	command.Flags().String("test-package", "", "Test application package")
	command.MarkFlagRequired("test-package")

	command.Flags().String("output", "", "Result artifacts folder path")
	command.MarkFlagRequired("output")

	command.Flags().Bool("verbose", false, "Use for additional output")
}

func DeclareCommonRunFlags(command *cobra.Command) {
	declareSharedFlags(command)
}

func DeclareIsolatedRunFlags(command *cobra.Command) {
	declareSharedFlags(command)
}
