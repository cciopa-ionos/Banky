package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	sumTr       int
	idTr        string
	description string

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
			auxValue := &Transaction{
				Id:          idTr,
				Amount:      sumTr,
				Description: description,
				Date:        time.Now(),
			}

			jsonFormating("./banky/operations.json", auxValue)

			// changing the deposit field in banky.json
			data, err := os.ReadFile("./banky/banky.json")
			check(err)

			var jsonArray []map[string]interface{}
			err = json.Unmarshal(data, &jsonArray)
			check(err)

			for _, obj := range jsonArray {
				idParse, exists := obj["Id"].(string)
				if exists && idParse == idTr {
					depositFloat, _ := obj["Deposit"].(float64)
					obj["Deposit"] = depositFloat + float64(sumTr)
				}
			}

			updatedData, err := json.MarshalIndent(jsonArray, "", "	")
			check(err)

			if err := os.WriteFile("./banky/banky.json", updatedData, 0666); err != nil {
				check(err)
			}

			if output == "json" {
				PrintTransactionJSON(auxValue)
			} else if output == "yaml" {
				PrintTransactionTable(auxValue)
			}
		},
	}
)

func init() {
	addCmd.Flags().StringVarP(&idTr, "id", "i", "", "id of account")
	addCmd.Flags().IntVarP(&sumTr, "sum", "s", 0, "sum of money added or retrieved")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "description of transaction")
	addCmd.Flags().StringVarP(&output, "output", "o", "table", "output format: table or json")
	viper.BindPFlag("output", addCmd.Flags().Lookup("output"))
	addCmd.MarkFlagRequired("id")
	addCmd.MarkFlagRequired("sum")
}
