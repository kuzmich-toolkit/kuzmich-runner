package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/kuzmich-toolkit/kuzmich-runner/arguments"
	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run instrumentation tests",
	Long:  `TODO: add full description`,

	Run: func(cmd *cobra.Command, args []string) {
		arguments := arguments.ParseCommonRunFlags(cmd)
		spew.Dump(arguments)
	},
}

func init() {
	rootCommand.AddCommand(runCommand)
	arguments.DeclareCommonRunFlags(runCommand)
}
