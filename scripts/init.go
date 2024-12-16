package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)
/*La c'est les import, bufio et os pour la manip des des fichiers, 
fmt pour l'affichage des messages
math/rand et time pour choisir un mot aléatoirement
strings pour les chaînes de carractéres*/

var WordsList []string
// On ajoute une variable WordsList ce sera pour stocker les les mots lus dans le fichier texte 

func LoadWords() {
	file, err := os.Open("DICTIONNAIRE/words.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			WordsList = append(WordsList, word)
		}
	}

	if err := scanner.Err(); err != nil {
		if err != nil {
			return
		}
	}
}
/*La func LoadWords sert à ouvrir le fichier txt qui est words.txt
os.Open("DICTIONNAIRE/wprds.txt") ouvre le fichier
if err!= nil c'est si y a un problème
defer file.Close() pour que le fichier soit toujours fermé
scanner := bufio.NewScanner(file) sert a lire un un fichier ligne par ligne
strings.TrimSpace , sa supprime les espaces*/


func PickWord() string {
	rand.Seed(time.Now().UnixNano())
	if len(WordsList) == 0 {
		return ""
	}
	return WordsList[rand.Intn(len(WordsList))]
}
/* du coup ici sa prends un mot aléatoirement à partir de WordsList
rand.Seed(time.Now().UnixNano()) permet d'initialiser le générateur de nombre aléatoire */



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
/*la fonction InitGame permet d'initialiser une nouvelle partie
WordsList qui sert toujours à choisir un mot aléatoire

if word == "" {
	fmt.Println("Erreur : Aucun mot disponible.")
	return
C'est pour si y a une erreur et y a un return pour relancer quand il y aura un mot valide

hidden := ""
for range word {
	hidden += "_"
}
Pour créer une version masqué du mot ça

LeJeu = Jeu{
	Word:    word,
	Found:   hidden,
	NbDeVie: 8,
	DejaMis: []rune{},
}
pour initialiser la partie, word le mot a deviner, found la version masqué du mot nbdevie le nombre de vie et dejamis un slice pour stocker les lettres déjà mise*/ 

func ProcessLetter(letter rune) {
	for _, l := range LeJeu.DejaMis {
		if l == letter {
			fmt.Println("Lettre déjà tentée :", string(letter))
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

/*LeJeu.Dejamis c'est pour les lettres déjà essayer
for _, l := range LeJeu.DejaMis c'est une boucle qui parcourt les lettres déjà essayer
et du coup si la lettre est déjà tenté sa mets le message suivant "Lettre déjà tenté"
et on ajoute return, la lettre sera pas ajouté une deuxième fois
LeJeu.DejaMis = append, append ajoute la lettre porposé

found := false
runes := []rune(LeJeu.Found)
for i, ch := range LeJeu.Word {
	if rune(ch) == letter {
		runes[i] = letter
		found = true
	}
}

Sa vérifie si la lettre est dans le mot a deviner et pour que sa cherche la lettre dans le mot on a fait
for i, ch := range LeJeu.Word , ch sa parcout les caractées


if !found {
	LeJeu.NbDeVie--
}

sa permet de réduire le nombre de vie*/