package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	accountCmd = &cobra.Command{
		Use:   "account",
		Short: "account commands",
		Long:  "account commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("GIVE ME ARGUMENTS!")
		},
	}
)

func init() {
	accountCmd.AddCommand(createCmd)
	accountCmd.AddCommand(welcomeCmd)
	accountCmd.AddCommand(listCmd)
}
