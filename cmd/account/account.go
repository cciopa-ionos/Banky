package account

import (
	"github.com/spf13/cobra"
)

var (
	AccountCmd = &cobra.Command{
		Use:   "account",
		Short: "account commands",
		Long:  "account commands",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

func init() {
	AccountCmd.AddCommand(createCmd)
	AccountCmd.AddCommand(welcomeCmd)
	AccountCmd.AddCommand(listCmd)
}
