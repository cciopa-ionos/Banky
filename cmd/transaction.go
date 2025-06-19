package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	TransactionCmd = &cobra.Command{
		Use:   "transaction",
		Short: "transaction commands",
		Long:  "transaction commands added to an account, or list the deposit amount",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("SMTH Works!")
		},
	}
)

func init() {
	RootCmd.AddCommand(TransactionCmd)
}

type person struct {
	id      string
	name    string
	deposit int
}

type transaction struct {
	sum         int
	date        string
	description string
}

func newPerson(id string, name string, deposit int) *person {
	return &person{
		id:      id,
		name:    name,
		deposit: deposit,
	}
}

func newTransaction(id string, description string) *transaction {
	// default date : today*
	return &transaction{
		sum:         0,
		description: description,
	}
}
