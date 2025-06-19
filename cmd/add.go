package cmd

import (
	"github.com/spf13/cobra"
)

var (
	deposit int
	amount  int
	from    string
	to      string

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add money to an account and subtract from another",
		Long:  "add money to an account",
		Run: func(cmd *cobra.Command, args []string) {

			//jsonTodo, err := json.Marshal(todo)
			//
			//if err != nil {
			//
			//	log.Fatalf("Error occurred during marshalling: %s", err.Error())
			//
			//}

		},
	}
)

func init() {
	addCmd.Flags().StringVarP(&Name, "sum", "s", "", "sum of money added to the account")
	TransactionCmd.AddCommand(addCmd)
}
