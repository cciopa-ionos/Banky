package core

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

var (
	letters          = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numberLetter int = 9
	Output       string
)

type Person struct {
	Id           string
	Name         string
	Transactions []Transaction
}

type Transaction struct {
	Amount      int
	Description string
	Date        time.Time
}

func RandSeq(n int) string {
	b := make([]rune, numberLetter)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func JsonFormating(jsonfile string, str interface{}) {
	var items []interface{}

	//check if banky.json is empty and then add []
	dataFile, _ := os.Stat(jsonfile)
	dataFileSize := dataFile.Size()

	if dataFileSize == 0 {
		if err := os.WriteFile(jsonfile, []byte("[]"), 0666); err != nil {
			fmt.Fprintf(os.Stderr, "Error adding [] to file %s: %v\n", jsonfile, err)
			os.Exit(1)
		}
	}

	fileData, err := os.ReadFile(jsonfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", jsonfile, err)
		os.Exit(1)
	}

	if len(fileData) == 0 {
		items = []interface{}{}
	} else {
		if err := json.Unmarshal(fileData, &items); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON from %s: %v\n", jsonfile, err)
			os.Exit(1)
		}
	}

	items = append(items, str)

	updatedData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON data: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(jsonfile, updatedData, 0666); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file %s: %v\n", jsonfile, err)
		os.Exit(1)
	}

}

func PrintPersonJSON(data Person) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}

func PrintPersonTable(data Person) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDEPOSIT")
	fmt.Fprintf(w, "%s\t%s\t%d\n", data.Id, data.Name)
	w.Flush()
}
