package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	name   string
	id     string
	scream bool

	welcomeCmd = &cobra.Command{
		Use:       "welcome",
		Short:     "cli for banky",
		Long:      "bankycli is a command line interface for banking transactions and operations ",
		ValidArgs: []string{"name", "id", "scream"},
		Run: func(cmd *cobra.Command, args []string) {
			var output string

			if name != "" && id != "" {
				fmt.Println("Error: Both name and id provided. Please provide only one.")
				return
			}

			if name != "" {
				output = name
			} else if id != "" {
				output = id
			}

			if scream {
				if output != "" {
					fmt.Printf("WELCOME %s TO BANKYCLI!\n", strings.ToUpper(output))
				} else {
					fmt.Println("WELCOME TO BANKYCLI!")
				}
			} else {
				if output != "" {
					fmt.Printf("Welcome %s to bankycli!\n", output)
				} else {
					fmt.Println("Welcome to bankycli!")
				}
			}
		},
	}
)

func init() {
	welcomeCmd.Flags().StringVarP(&name, "name", "n", "", "name of the account")
	welcomeCmd.Flags().StringVarP(&id, "id", "i", "", "id of the account")
	welcomeCmd.Flags().BoolVar(&scream, "scream", false, "scream account")
	rootCmd.AddCommand(welcomeCmd)
}
