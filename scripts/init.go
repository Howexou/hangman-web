package hangman

import (
	"fmt"
	"math/rand"
	"time"
)

var WordsList []string 

func LoadWords() {
	WordsList = []string{"chien", "soleil", "porte", "chat"}
}

func PickWord() string {
	rand.Seed(time.Now().UnixNano())
	if len(WordsList) == 0 {
		return ""
	}
	return WordsList[rand.Intn(len(WordsList))]
}

func InitGame() {
	word := PickWord()
	if word == "" {
		fmt.Println("Erreur : Aucun mot disponible.")
		return
	}

	hidden := ""
	for range word {
		hidden += "_"
	}

	LeJeu = Jeu{
		Word:    word,
		Found:   hidden,
		NbDeVie: 8,
		DejaMis: []rune{},
	}
}

func ProcessLetter(letter rune) {
	for _, l := range LeJeu.DejaMis {
		if l == letter {
			return
		}
	}

	LeJeu.DejaMis = append(LeJeu.DejaMis, letter)

	found := false
	runes := []rune(LeJeu.Found)
	for i, ch := range LeJeu.Word {
		if rune(ch) == letter {
			runes[i] = letter
			found = true
		}
	}

	LeJeu.Found = string(runes)

	if !found {
		LeJeu.NbDeVie--
	}
}
