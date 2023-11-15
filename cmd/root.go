package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "awm",
	Short: "AWM is a command-line to parse awm transaction.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to AWM!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(txCmd)
	rootCmd.AddCommand(accessListCmd)
	rootCmd.AddCommand(warpUnsignedMsgCmd)
}
