package transaction

import (
	"bankycli/internal/core"
	"encoding/json"
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
		Run: func(cmd *cobra.Command, args []string) {
			if sumTr == 0 {
				print("suma is required!\n")
			} else if idTr == "" {
				print("id is required!\n")
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
			core.Check(err)

			var jsonArray []map[string]interface{}
			err = json.Unmarshal(data, &jsonArray)
			core.Check(err)

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
			core.Check(err)

			if err := os.WriteFile("./banky/banky.json", updatedData, 0666); err != nil {
				core.Check(err)
			}

			if output == "json" {
				core.PrintTransactionJSON(auxValue)
			} else if output == "table" {
				core.PrintTransactionTable(auxValue)
			}
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
