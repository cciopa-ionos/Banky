package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	name string

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "create account",
		Long:  "create new account",
		Run: func(cmd *cobra.Command, args []string) {

			// Create a Person struct
			if name != "" {
				fmt.Printf("Created account for %s. \nWelcome to bankycli!\n", name)
			}
			auxValue := &Person{Id: randSeq(7), Name: name, Deposit: 0}

			jsonFormating("./banky/banky.json", auxValue)
		},
	}
)

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "name of the owner")
	createCmd.MarkFlagRequired("name")
}
