package main

import (
	h "main/scripts"
	"fmt"
)

func main() {
	h.LoadWords()
if len(h.WordsList) == 0 {
	fmt.Println("Erreur : Aucun mot disponible dans le dictionnaire.")
	return
}
	h.InitGame()
	h.HandleRequests()
}
