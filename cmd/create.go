package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	name string

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "create account",
		Long:  "create new account",
		Run: func(cmd *cobra.Command, args []string) {

			// Create a Person struct
			if name != "" {
				fmt.Printf("Created account for %s. \nWelcome to bankycli!\n", name)
			}
			auxValue := &Person{Id: randSeq(7), Name: name, Deposit: 0}

			//Marshal the respective struct
			jsonData, err := json.MarshalIndent(auxValue, "", "	")
			if err != nil {
				fmt.Printf("could not marshal json: %s\n", err)
				return
			}

			//TODO: Delete
			fmt.Printf("json data: %s\n", jsonData)

			//	Check if file is empty and add []
			stats, err := os.Stat("./banky/banky.json")
			check(err)
			if stats.Size() == 0 {
				initial := []byte("[\n")
				initial = append(initial, jsonData...)
				initial = append(initial, '\n', ']')
				err = os.WriteFile("./banky/banky.json", initial, 0644)
				check(err)
			}

			stats, err = os.Stat("./banky/banky.json")
			existingData, err := os.ReadFile("./banky/banky.json")
			check(err)

			if stats.Size() != 4 {
				// not first json data
				insertion := append([]byte("\n"), jsonData...)
				insertion = append(insertion, ',', '\n')

				// always add the [ back after insertion
				newContent := append([]byte{'['}, insertion...)
				newContent = append(newContent, existingData[1:]...)
				err = os.WriteFile("./banky/banky.json", newContent, 0666)
				check(err)
			} else {
				// first json data
				jsonData = append(jsonData, existingData[1:]...)
				err = os.WriteFile("./banky/banky.json", jsonData, 0666)
				check(err)
			}

		},
	}
)

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "name of the owner")
	accountCmd.AddCommand(createCmd)
}
