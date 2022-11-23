package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func main() {
	var echoTimes int

	var cmdPrint = &cobra.Command{
		Use:   "Print",
		Short: "print anything to the screen",
		Long:  "print",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, ""))
		},
	}

	var cmdTimes = &cobra.Command{
		Use:   "times",
		Short: "echo anything to the screen",
		Long:  "echo",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, ""))
			}
		},
	}

	var cmdEcho = &cobra.Command{
		Use:   "echo",
		Short: "echo anything to the screen",
		Long:  "echo",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, ""))
		},
	}

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "time to echo th input")
	var rootCmd = &cobra.Command{
		Use: "app",
	}
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	cmdEcho.AddCommand(cmdTimes)
	rootCmd.Execute()
}
