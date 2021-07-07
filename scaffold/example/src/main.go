package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env, %v", err)
	}
}

func main() {
	log.Println("Ehlo World")
}
