package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run instrumentation tests",
	Long:  `TODO: add full description`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
	},
}

func init() {
	rootCommand.AddCommand(runCommand)
}
