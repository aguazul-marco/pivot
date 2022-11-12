package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aguazul-marco/pivot/marvel"
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

	characters, err := client.GetCharacters(5)
	if err != nil {
		log.Fatal(err)
	}

	for _, character := range characters {
		fmt.Printf("Character: %v\nDescription: %v\n\n", character.Name, character.Description)
	}

}
