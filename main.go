/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bankycli/cmd"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
	cmd.Execute()
}
