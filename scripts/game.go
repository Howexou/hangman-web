package hangman

import (
	"html/template"
	"net/http"
	"strconv"
)

// Initialiser le jeu
func initGame() {
	word = PickWord()          // Mot à deviner
	hiddenWord = ""            // Réinitialise le mot caché
	for range word {           // Crée "_" pour chaque lettre
		hiddenWord += "_"
	}
}

func Play(w http.ResponseWriter, r *http.Request) {
	// On initialise
	if word == "" {
		initGame()
	}

	// Gestion de la lettre devinée
	if r.Method == "POST" {
		r.ParseForm()
		guess := r.FormValue("guess") // Récupère la lettre qui a été guess
		corectguess := false

		// Mise à jour du mot caché
		newHiddenWord := ""
		for i, char := range word {
			if string(char) == guess { // Lettre correcte
				newHiddenWord += string(char)
				corectguess = true
			} else {
				newHiddenWord += string(hiddenWord[i]) // Conserve les lettres déjà trouvées
			}
		}
		if !corectguess {
			lives -= 1
		}
		
		hiddenWord = newHiddenWord
	}

	

	// Structure des données envoyées au template
	data := struct {
		HiddenWord string
		PhasePendu string
		lives int

	} {
		HiddenWord: hiddenWord,
		PhasePendu: "/static/hangman-game-images/images/hangman-" + strconv.Itoa(6 - lives) + ".svg",
		lives: lives,
	}

	// Charger et afficher le template
	t, err := template.ParseFiles("./templates/Play.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)

	


}
