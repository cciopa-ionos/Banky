package transaction

import (
	"github.com/spf13/cobra"
)

var (
	TransactionCmd = &cobra.Command{
		Use:   "transaction",
		Short: "transaction commands",
		Long:  "transaction commands",
	}
)

func init() {
	TransactionCmd.AddCommand(addCmd)
	TransactionCmd.AddCommand(listTraCmd)
}
