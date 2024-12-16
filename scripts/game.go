package hangman

import (
	"html/template"
	"net/http"
)

// Initialiser le jeu
func initGame() {
	word = PickWord()          // Mot à deviner
	hiddenWord = ""            // Réinitialise le mot caché
	for range word {           // Crée "_" pour chaque lettre
		hiddenWord += "_"
	}
}

func Play(w http.ResponseWriter, r *http.Request) { // Fonction de jeu
	// On initialise
	if word == "" {
		initGame()
	}

	// Gestion de la lettre devinée
	if r.Method == "POST" {
		r.ParseForm()
		guess := r.FormValue("guess") // Récupère la lettre qui a été guess

		// Mise à jour du mot caché
		newHiddenWord := ""
		for i, char := range word {
			if string(char) == guess {
				newHiddenWord += string(char)
			} else {
				newHiddenWord += string(hiddenWord[i]) // Conserve les lettres déjà trouvées
			}
		}
		hiddenWord = newHiddenWord
	}

	// Structure des données envoyées au template
	data := struct {
		HiddenWord string
	}{
		HiddenWord: hiddenWord,
	}

	// Charger et afficher le template
	t, err := template.ParseFiles("./templates/Play.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
