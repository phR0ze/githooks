package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Prepend version to every commit message",
		Long: `Read the project version from the target file using the target format
then prepend the version to the commit message.

Examples:

  githooks blah bhlah
`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Prepending version to commit message")
		},
	}

	return cmd
}
