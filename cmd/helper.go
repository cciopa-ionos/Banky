package cmd

import (
	"encoding/json"
	"math/rand"
	"os"
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

	//Marshal the respective struct
	jsonData, err := json.MarshalIndent(str, "", "	")
	check(err)

	//	Check if file is empty and add []
	stats, err := os.Stat(jsonfile)
	check(err)
	existingData, err := os.ReadFile(jsonfile)
	check(err)
	if stats.Size() == 0 {
		initial := []byte("[\n")
		initial = append(initial, jsonData...)
		initial = append(initial, '\n', ']')
		err = os.WriteFile(jsonfile, initial, 0644)
		check(err)
	} else {
		// not first json data
		insertion := append([]byte("\n"), jsonData...)
		insertion = append(insertion, ',', '\n')

		// always add the [ back after insertion
		newContent := append([]byte{'['}, insertion...)
		newContent = append(newContent, existingData[1:]...)
		err = os.WriteFile(jsonfile, newContent, 0666)
		check(err)
	}
}
