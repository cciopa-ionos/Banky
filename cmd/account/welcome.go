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
		RunE: func(cmd *cobra.Command, args []string) error {
			if scream == "yes" && user != "" {
				fmt.Printf("HELLO %v\n", strings.ToUpper(user))
				return nil
			} else if scream == "no" && user != "" {
				fmt.Printf("Hello %v\n", user)
				return nil
			}

			cfg := core.LoadConfig()

			data, err := os.ReadFile(cfg.BankyPath)
			if err != nil {
				return fmt.Errorf("error reading file %s: %w", cfg.BankyPath, err)
			}

			var jsonArray []map[string]interface{}
			if err := json.Unmarshal(data, &jsonArray); err != nil {
				return fmt.Errorf("error parsing JSON from %s: %w", cfg.BankyPath, err)
			}

			for _, obj := range jsonArray {
				if obj["Id"] == id {
					nameVal, _ := obj["Name"]
					nameStr, _ := nameVal.(string)
					if scream == "yes" {
						fmt.Printf("HELLO %v\n", strings.ToUpper(nameStr))
					} else if scream == "no" {
						fmt.Printf("Hello %v\n", nameStr)
					}
					return nil
				}
			}
			fmt.Fprintf(os.Stderr, "User with ID %s not found.\n", id)
			return nil
		},
	}
)

func init() {
	welcomeCmd.Flags().StringVarP(&user, "user", "u", "", "name of the account")
	welcomeCmd.Flags().StringVarP(&id, "id", "i", "", "id of the account")
	welcomeCmd.Flags().StringVarP(&scream, "scream", "s", "no", "scream flag")
}
