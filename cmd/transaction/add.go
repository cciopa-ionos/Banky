package transaction

import (
	"bankycli/internal/core"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	sumTr       int
	idTr        string
	description string
	output      string

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add sum of money",
		Long:  "add sum of money",
		RunE: func(cmd *cobra.Command, args []string) error {
			if sumTr == 0 {
				fmt.Fprintf(os.Stderr, "suma is required!\n")
				os.Exit(1)
			} else if idTr == "" {
				fmt.Fprintf(os.Stderr, "suma is required!\n")
				os.Exit(1)
			}

			//adding operation to operations.json
			auxValue := &core.Transaction{
				Amount:      sumTr,
				Description: description,
				Date:        time.Now(),
			}

			// Load config
			cfg := core.LoadConfig()

			// Read update banky.json using config
			data, err := os.ReadFile(cfg.BankyPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading file %s: %v\n", cfg.BankyPath, err)
				os.Exit(1)
			}

			var jsonArray []map[string]interface{}
			if err := json.Unmarshal(data, &jsonArray); err != nil {
				fmt.Fprintf(os.Stderr, "error parsing JSON: %v\n", err)
				os.Exit(1)
			}

			for _, obj := range jsonArray {
				idParse, exists := obj["Id"].(string)
				if exists && idParse == idTr {
					depositFloat, _ := obj["Deposit"].(float64)
					obj["Deposit"] = depositFloat + float64(sumTr)

					arrayTransactions, _ := obj["Transactions"].([]interface{})
					obj["Transactions"] = append(arrayTransactions, auxValue)
				}
			}

			updatedData, err := json.MarshalIndent(jsonArray, "", "	")
			if err != nil {
				fmt.Fprintf(os.Stderr, "error marshalling JSON: %v\n", err)
				os.Exit(1)
			}

			if err := os.WriteFile("./banky/banky.json", updatedData, 0666); err != nil {
				fmt.Fprintf(os.Stderr, "error writing to file %s: %v\n", cfg.BankyPath, err)
				os.Exit(1)
			}

			return nil
		},
	}
)

func init() {
	addCmd.Flags().StringVarP(&idTr, "id", "i", "", "id of account")
	addCmd.Flags().IntVarP(&sumTr, "sum", "s", 0, "sum of money added or retrieved")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "description of transaction")
	addCmd.Flags().StringVarP(&output, "output", "o", "", "output format: table or json")
	addCmd.MarkFlagRequired("id")
	addCmd.MarkFlagRequired("sum")
}
