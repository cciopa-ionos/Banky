package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	transactionCmd = &cobra.Command{
		Use:   "transaction",
		Short: "transaction commands",
		Long:  "transaction commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("GIVE ME ARGUMENTS!")
		},
	}
)

func init() {
	transactionCmd.AddCommand(addCmd)
	transactionCmd.AddCommand(listTraCmd)
}
