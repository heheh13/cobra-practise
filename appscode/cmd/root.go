package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	persistentFlag bool
	localFlag      bool
	times          int
	echoFlag       bool
	// version        bool

	rootCmd = &cobra.Command{
		Use:   "CLI",
		Short: "a example of cobra",
		Long: `cobra is used to develop cli base application.
				which was developed by spf13`,
		Version: "v1.0",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello from the root command", persistentFlag)
			cmd.SetVersionTemplate("v1.1")

			// fmt.Println("v3.o")
		},
	}

	echoCmd = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "prints given string to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if persistentFlag {
				fmt.Println("persistant echo")
			}
			fmt.Println("> : " + strings.Join(args, " "))
		},
	}

	timesCmd = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "number of time you want to execute a command",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if times == 0 {
				return errors.New("times cant be zero")
			}
			for i := 0; i < times; i++ {
				fmt.Println("> : " + strings.Join(args, " "))
			}
			return nil
		},
	}
)

func init() {

	rootCmd.PersistentFlags().BoolVarP(&persistentFlag, "persist", "p", false, "a persistent flag")
	rootCmd.Flags().BoolVarP(&localFlag, "local-flag", "l", false, "a local flag")
	// rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "prints the version")
	echoCmd.PersistentFlags().BoolVarP(&echoFlag, "echoFlag", "e", false, "echo persistant flag")
	timesCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times")

	rootCmd.AddCommand(echoCmd)
	echoCmd.AddCommand(timesCmd)
}

//Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}
