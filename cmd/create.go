package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

			if output == "json" {
				PrintPersonJSON(*auxValue)
			} else if output == "yaml" {
				PrintPersonTable(*auxValue)
			}
		},
	}
)

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "name of the owner")
	createCmd.Flags().StringVarP(&name, "output", "o", "table", "output format: table or json")
	viper.BindPFlag("output", createCmd.Flags().Lookup("output"))
	createCmd.MarkFlagRequired("name")

}
