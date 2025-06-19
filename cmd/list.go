package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list money to an account and subtract from another",
		Long:  "list money to an account",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("WELCOME TO LIST!")
			//jsonTodo, err := json.Marshal(todo)

			//if err != nil {
			//
			//	log.Fatalf("Error occurred during marshalling: %s", err.Error())
			//
			//}

		},
	}
)

func init() {
	addCmd.Flags().StringVarP(&Name, "list", "l", "", "list the sum of money added to the account")
	TransactionCmd.AddCommand(listCmd)
}
