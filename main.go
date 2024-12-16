package main

import (
	"fmt"
	h "main/scripts"
)

func main() {
	
	h.LoadWords()

	if len(h.WordsList) == 0 {
		fmt.Println("Erreur : Aucun mot trouvé dans le fichier words.txt.")
		return
	}

	h.InitGame()

	h.HandleRequests()
}
