package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command executed when the app is called with no arguments
var rootCmd = &cobra.Command{
	Use:  "githooks",
	Long: "Set of githooks to increment versions, update copyrights etc...",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(
		newVersionCmd(),
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
