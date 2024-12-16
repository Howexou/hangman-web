package main

import (
	"fmt"
	h "main/scripts"
)

func main() {
	// Charger les mots depuis le fichier words.txt
	h.LoadWords()

	// Vérifie si des mots sont disponibles
	if len(h.WordsList) == 0 {
		fmt.Println("Erreur : Aucun mot trouvé dans le fichier words.txt.")
		return
	}

	// Démarrer une partie
	h.InitGame()

	// Lancer le serveur web
	h.HandleRequests()
}
