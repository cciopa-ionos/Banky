package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	Name   string
	id     string
	scream bool

	welcomeCmd = &cobra.Command{
		Use:   "welcome",
		Short: "welcome print for this bankycli",
		Long:  "bankycli is a command line interface for banking transactions and operations",
		Run: func(cmd *cobra.Command, args []string) {
			var output string

			if Name != "" && id != "" {
				fmt.Println("Error: Both name and id provided. Please provide only one.")
				return
			}

			if Name != "" {
				output = Name
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
	welcomeCmd.Flags().StringVarP(&Name, "name", "n", "", "name of the account")
	welcomeCmd.Flags().StringVarP(&id, "id", "i", "", "id of the account")
	welcomeCmd.Flags().BoolVar(&scream, "scream", false, "scream account")
	RootCmd.AddCommand(welcomeCmd)
}
