package cmd

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

func randSeq(n int) string {
	b := make([]rune, numberLetter)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func check(err error) {
	if err != nil {
		print(err)
	}
}

func jsonFormating(jsonfile string, str interface{}) {

	////Marshal the respective struct
	//jsonData, err := json.MarshalIndent(str, "", "	")
	//check(err)
	//
	////	Check if file is empty and add []
	//stats, err := os.Stat(jsonfile)
	//check(err)
	//existingData, err := os.ReadFile(jsonfile)
	//check(err)
	//if stats.Size() == 0 {
	//	initial := []byte("[\n")
	//	initial = append(initial, jsonData...)
	//	initial = append(initial, '\n', ']')
	//	err = os.WriteFile(jsonfile, initial, 0666)
	//	check(err)
	//} else {
	//	// not first json data
	//	insertion := append([]byte("\n"), jsonData...)
	//	insertion = append(insertion, ',', '\n')
	//
	//	// always add the [ back after insertion
	//	newContent := append([]byte{'['}, insertion...)
	//	newContent = append(newContent, existingData[1:]...)
	//	err = os.WriteFile(jsonfile, newContent, 0666)
	//	check(err)
	//}
	var items []interface{}

	fileData, err := os.ReadFile(jsonfile)
	if err != nil {
		// If file is empty, start with empty slice
		if os.IsNotExist(err) || len(fileData) == 0 {
			items = []interface{}{}
		} else {
			check(err)
		}
	} else {
		if err := json.Unmarshal(fileData, &items); err != nil {
			check(err)
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
	fmt.Fprintf(w, "%s\t %s\t %d\t %d\n", data.Id, data.Amount, data.Description, data.Date)
}
