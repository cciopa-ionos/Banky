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
	Id      string
	Name    string
	Deposit int
}

type Transaction struct {
	Id          string
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

func Check(err error) {
	if err != nil {
		print(err)
	}
}

func JsonFormating(jsonfile string, str interface{}) {
	var items []interface{}

	fileData, err := os.ReadFile(jsonfile)
	if err != nil {
		// If file is empty, start with empty slice
		if os.IsNotExist(err) || len(fileData) == 0 {
			items = []interface{}{}
		} else {
			Check(err)
		}
	} else {
		if err := json.Unmarshal(fileData, &items); err != nil {
			Check(err)
		}
	}

	items = append(items, str)

	updatedData, _ := json.MarshalIndent(items, "", "  ")

	os.WriteFile(jsonfile, updatedData, 0666)
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
	fmt.Fprintf(w, "%s\t %s\t %d\n", data.Id, data.Name, data.Deposit)
	w.Flush()
}

func PrintTransactionJSON(data *Transaction) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}

func PrintTransactionTable(data *Transaction) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tAMOUNT\tDESCRIPTION\tDATE")
	formattedDate := data.Date.Format("2022-01-02 15:04")
	fmt.Fprintf(w, "%s\t %d\t %s\t %s\n", data.Id, data.Amount, data.Description, formattedDate)
	w.Flush()
}
