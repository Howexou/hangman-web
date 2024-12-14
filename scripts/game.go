package hangman

import (
	"html/template"
	"net/http"
)

func Play(w http.ResponseWriter, r *http.Request) {
	word := PickWord()
	hiddenWord := ""
	
	for range word {
        hiddenWord += "_"

	if r.Method == "POST" {
		r.ParseForm()
		guess := r.FormValue("guess")

		// Mettre à jour les "_" si la lettre devinée est correcte
		newHiddenWord := ""
		for i, char := range word {
			if string(char) == guess {
				newHiddenWord += string(char)
			} else {
				newHiddenWord += string(hiddenWord[i])
			}
		}
		hiddenWord = newHiddenWord
	}

	data := struct {
		Word       string
		HiddenWord string
	}{
		Word:       word,
		HiddenWord: hiddenWord,
	}

	t, err := template.ParseFiles("./templates/Play.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
