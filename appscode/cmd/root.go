package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	persitantflag bool
	rootCmd       = &cobra.Command{
		Use:   "example",
		Short: "a exaple of cobra",
		Long: `cobra is used to devlop cli base application.
				which was deleloped by spf13`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello from the root command")
		},
	}
	echoCmd = &cobra.Command{
		Use:   "echo [stirng to echo]",
		Short: "prints given stirng to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("echo: " + strings.Join(args, " "))
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persitantflag, "persist", "p", false, "a persistant flag")
	rootCmd.AddCommand(echoCmd)
}

//Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}
