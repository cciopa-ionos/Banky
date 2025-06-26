/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bankycli/cmd/account"
	"bankycli/cmd/transaction"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	bankycliCmd = &cobra.Command{
		Use:              "bankycli",
		Short:            "cli for banking",
		Long:             "bankycli is a command line interface for banking transactions and operations ",
		TraverseChildren: true,
	}
)

func Execute() {
	if err := bankycliCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	bankycliCmd.AddCommand(account.AccountCmd)
	bankycliCmd.AddCommand(transaction.TransactionCmd)
}
