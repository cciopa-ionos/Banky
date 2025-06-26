package account

import (
	"bankycli/internal/core"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	name   string
	output string

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "create account",
		Long:  "create new account",
		Run: func(cmd *cobra.Command, args []string) {

			// Create a Person struct
			if name != "" {
				fmt.Printf("Created account for %s. \nWelcome to bankycli!\n", name)
			}
			auxValue := &core.Person{Id: core.RandSeq(7), Name: name, Deposit: 0}

			core.JsonFormating("./banky/banky.json", auxValue)

			if output == "json" {
				core.PrintPersonJSON(*auxValue)
			} else if output == "table" {
				core.PrintPersonTable(*auxValue)
			}
		},
	}
)

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "name of the owner")
	createCmd.Flags().StringVarP(&output, "output", "o", "", "output format: table or json")
	createCmd.MarkFlagRequired("name")

}
