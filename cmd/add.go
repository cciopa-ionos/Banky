package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
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

			jsonTodo, err := json.Marshal(todo)

			if err != nil {

				log.Fatalf("Error occurred during marshalling: %s", err.Error())

			}

		},
	}
)

func init() {
	welcomeCmd.Flags().StringVarP(&name, "receiver", "n", "", "name of the receiver")
	welcomeCmd.Flags().StringVarP(&id, "initiator", "i", "", "name of the initiator")
	rootCmd.AddCommand(addCmd)
}
