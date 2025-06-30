package transaction

import (
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
		RunE: func(cmd *cobra.Command, args []string) error {

			data, err := os.ReadFile("./banky/banky.json")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading file ./banky/banky.json: %v\n", err)
				os.Exit(1)
			}

			var jsonArray []map[string]interface{}
			if err := json.Unmarshal(data, &jsonArray); err != nil {
				fmt.Fprintf(os.Stderr, "error parsing JSON from ./banky/banky.json: %v\n", err)
				os.Exit(1)
			}

			found := false
			fmt.Printf("The account with ID: %v\n\n", idList)
			for _, obj := range jsonArray {
				idParse, exists := obj["Id"].(string)
				if exists && idParse == idList {
					for _, t := range obj["Transactions"].([]interface{}) {
						transaction, _ := t.(map[string]interface{})
						fmt.Printf("Amount: %v\n", transaction["Amount"].(float64))
						fmt.Printf("Date: %v\n", transaction["Date"].(string))
						fmt.Printf("Description: %v\n\n", transaction["Description"].(string))
						found = true
						break
					}
				}
			}
			if !found {
				fmt.Fprintf(os.Stderr, "account with ID '%s' not found\n", idList)
				os.Exit(1)
			}

			return nil
		},
	}
)

func init() {
	listTraCmd.Flags().StringVarP(&idList, "id", "i", "", "id of the account you want to list")
	listTraCmd.MarkFlagRequired("id")
}
