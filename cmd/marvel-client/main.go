package main

import (
	"log"
	"os"

	"github.com/aguazul-marco/pivot/marvel"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pubKey := os.Getenv("MARVEL_PUBLIC_KEY")
	privKey := os.Getenv("MARVEL_PRIVATE_KEY")

	client := marvel.NewClient(pubKey, privKey)

	character, err := client.GetCharacter(5)
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(character)
}
