package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var welcomeCmd = &cobra.Command{
	Use:   "welcome",
	Short: "cli for banky",
	Long:  "bankycli is a command line interface for banking transactions and operations ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Welcome to bankycli!\n")
	},
}

func init() {
	rootCmd.AddCommand(welcomeCmd)
}
