package hangman

import (
	"bufio"
	"fmt"
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
	r := bufio.NewReader(f)

	// Lire ligne par ligne
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		WordsList = append(WordsList, line)
	}

	// Fermer le fichier
	defer f.Close()
}

// Choix de mot aléatoire dans words.txt
func PickWord() string {
	Word = WordsList[rand.Intn(len(WordsList))]
	return Word
}
