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

			data, err := os.ReadFile("./banky/operations.json")
			core.Check(err)

			var jsonArray []map[string]interface{}
			err = json.Unmarshal(data, &jsonArray)
			core.Check(err)

			fmt.Printf("The account with ID: %v\n\n", idList)
			for _, obj := range jsonArray {
				idParse, exists := obj["Id"].(string)
				if exists && idParse == idList {
					fmt.Printf("Amount transfered: %v\n", obj["Amount"].(float64))
					fmt.Printf("Description: %v\n", obj["Description"].(string))
					fmt.Printf("Date of the transaction: %v\n\n", obj["Date"].(string))
				}
			}
		},
	}
)

func init() {
	listTraCmd.Flags().StringVarP(&idList, "id", "i", "", "id of the account you want to list")
	listTraCmd.MarkFlagRequired("id")
}
