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
	user   string
	id     string
	scream string

	welcomeCmd = &cobra.Command{
		Use:   "welcome",
		Short: "welcome print",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if scream == "yes" && user != "" {
				fmt.Printf("HELLO %v\n", strings.ToUpper(user))
				return
			} else if scream == "no" && user != "" {
				fmt.Printf("Hello %v\n", user)
				return
			}

			cfg := core.LoadConfig()

			data, err := os.ReadFile(cfg.BankyPath)
			core.Check(err)

			var jsonArray []map[string]interface{}
			err = json.Unmarshal(data, &jsonArray)
			core.Check(err)

			for _, obj := range jsonArray {
				//fmt.Printf("Hello %v\n", obj["Id"])
				//for key, val := range obj {
				//	if key == "Id" {
				//		valStr, ok := val.(string)
				//		if !ok {
				//			continue
				//		}
				if obj["Id"] == id {
					nameVal, _ := obj["Name"]
					nameStr, _ := nameVal.(string)
					if scream == "yes" {
						fmt.Printf("HELLO %v\n", strings.ToUpper(nameStr))
					} else if scream == "no" {
						fmt.Printf("Hello %v\n", nameStr)
					}
					return
				}
				//}
				//}
			}
		},
	}
)

func init() {
	welcomeCmd.Flags().StringVarP(&user, "user", "u", "", "name of the account")
	welcomeCmd.Flags().StringVarP(&id, "id", "i", "", "id of the account")
	welcomeCmd.Flags().StringVarP(&scream, "scream", "s", "no", "scream flag")
}
