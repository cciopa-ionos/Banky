package account

import (
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
			bankyPath := os.Getenv("BANKY_PATH")
			if bankyPath == "" {
				bankyPath = "./banky/banky.json"
			}
			data, err := os.ReadFile(bankyPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", bankyPath, err)
				os.Exit(1)
			}

			// Parse JSON as slice of maps
			var jsonArray []map[string]interface{}
			if err := json.Unmarshal(data, &jsonArray); err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing JSON from %s: %v\n", bankyPath, err)
				os.Exit(1)
			}

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
