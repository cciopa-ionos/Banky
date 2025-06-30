package account

import (
	"bankycli/internal/core"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var (
	name   string
	output string

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "create account",
		Long:  "create new account",
		Run: func(cmd *cobra.Command, args []string) {

			if name != "" {
				fmt.Printf("Created account for %s. \nWelcome to bankycli!\n", name)
			}
			auxValueTrans := &core.Transaction{
				Amount:      0,
				Description: "Creating Account",
				Date:        time.Now(),
			}

			auxValue := &core.Person{Id: core.RandSeq(7), Name: name, Transactions: []core.Transaction{*auxValueTrans}}

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
