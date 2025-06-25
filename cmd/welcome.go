package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
	user   string
	id     string
	scream string

	welcomeCmd = &cobra.Command{
		Use:   "welcome",
		Short: "welcome print",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			data, err := os.ReadFile("./banky/banky.json")
			check(err)

			var jsonArray []map[string]interface{}
			err = json.Unmarshal(data, &jsonArray)
			check(err)

			for _, obj := range jsonArray {
				for key, val := range obj {
					if key == "Id" {
						valStr, ok := val.(string)
						if !ok {
							continue
						}
						if valStr == id {
							nameVal, _ := obj["Name"]
							nameStr, _ := nameVal.(string)
							fmt.Println("DEBUG: scream flag =", scream)
							if scream == "yes" {
								fmt.Printf("HELLO %v\n", strings.ToUpper(nameStr))
							} else if scream == "no" {
								fmt.Printf("Hello %v\n", nameStr)
							}
							return
						}
					}
				}
			}
		},
	}
)

func init() {
	welcomeCmd.Flags().StringVarP(&user, "user", "u", "", "name of the account")
	welcomeCmd.Flags().StringVarP(&id, "id", "i", "", "id of the account")
	welcomeCmd.Flags().StringVarP(&scream, "scream", "s", "no", "scream flag")
}
