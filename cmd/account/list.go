package account

import (
	"bankycli/internal/core"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list users name",
		Long:  "list a list of users name",
		Run: func(cmd *cobra.Command, args []string) {
			//data, err := os.ReadFile("./banky/banky.json")
			//core.Check(err)
			bankyPath := os.Getenv("BANKY_PATH")
			if bankyPath == "" {
				bankyPath = "./banky/banky.json"
			}
			data, err := os.ReadFile(bankyPath)

			// Parse JSON as slice of maps
			var jsonArray []map[string]interface{}
			err = json.Unmarshal(data, &jsonArray)
			core.Check(err)

			//Print Names
			fmt.Printf("Names:")
			for _, obj := range jsonArray {
				for key, val := range obj {
					if strings.HasPrefix(key, "Name") {
						fmt.Printf("\n %v", val)
					}
				}
			}
			fmt.Printf("\n\n")

		},
	}
)
