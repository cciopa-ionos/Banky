package cmd

import (
	"math/rand"
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
