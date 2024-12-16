package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// Importe les mots dans un tableau de string
func Words() {
	// Ouvre le fichier
	f, err := os.Open("DICTIONNAIRE/words.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		WordsList = append(WordsList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Fermer le fichier
	defer f.Close()
}

// Choix de mot aléatoire dans words.txt
func PickWord() string {
	Word = WordsList[rand.Intn(len(WordsList))]
	return Word
}
