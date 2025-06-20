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
	scream bool

	welcomeCmd = &cobra.Command{
		Use:   "welcome",
		Short: "welcome print",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if user != "" && scream != true {
				fmt.Printf("Hello %v\n", user)
				return
			} else if scream == true {
				fmt.Printf("Hello %v\n", strings.ToUpper(user))
				return
			}

			data, err := os.ReadFile("./banky/banky.json")
			check(err)

			// Parse JSON as slice of maps
			var jsonArray []map[string]interface{}
			err = json.Unmarshal(data, &jsonArray)
			check(err)

			aux := 0
			if id != "" {
				for _, obj := range jsonArray {
					for key, val := range obj {
						if key == "Id" && val == id {
							aux = 1
						}
						if aux == 1 && key == "Name" {
							aux = 0
							if scream {

								fmt.Printf("HELLO %v\n", strings.ToUpper(val.(string)))
								return
							}
							fmt.Printf("Hello %v\n", val)
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
	welcomeCmd.Flags().BoolVar(&scream, "scream", false, "scream account")
	accountCmd.AddCommand(welcomeCmd)
}
