package transaction

import (
	"bankycli/internal/core"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	idList string

	listTraCmd = &cobra.Command{
		Use:   "list",
		Short: "List transactions for a user account",
		Long:  "List transactions for a user account",
		Run: func(cmd *cobra.Command, args []string) {

			data, err := os.ReadFile("./banky/banky.json")
			core.Check(err)

			var jsonArray []map[string]interface{}
			err = json.Unmarshal(data, &jsonArray)
			core.Check(err)

			fmt.Printf("The account with ID: %v\n\n", idList)
			for _, obj := range jsonArray {
				idParse, exists := obj["Id"].(string)
				if exists && idParse == idList {
					for _, t := range obj["Transactions"].([]interface{}) {
						transaction, _ := t.(map[string]interface{})
						fmt.Printf("Amount: %v\n", transaction["Amount"].(float64))
						fmt.Printf("Date: %v\n", transaction["Date"].(string))
						fmt.Printf("Description: %v\n\n", transaction["Description"].(string))
					}
					return
				}
			}
		},
	}
)

func init() {
	listTraCmd.Flags().StringVarP(&idList, "id", "i", "", "id of the account you want to list")
	listTraCmd.MarkFlagRequired("id")
}
