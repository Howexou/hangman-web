package hangman

import (
	"html/template"
	"net/http"
	"strconv"
)

// Initialiser le jeu
func initGame() {
	word = PickWord() 	// Mot à deviner
	hiddenWord = ""  	 // Réinitialise le mot caché
	for range word {  	// Crée "_" pour chaque lettre
		hiddenWord += "_"
	}
	lives = 6
}

func ResetGame(w http.ResponseWriter, r *http.Request) {
	initGame()
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}

func Play(w http.ResponseWriter, r *http.Request) { // Fonction de jeu
	win := false

	if word == "" { // On initialise
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
			if string(char) == guess {
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

		// Vérification si le joueur a gagné
		win = CheckWin(hiddenWord)

		if lives <= 0 {
			RenderTemplate(w, "templates/Lose.html", nil)
			return

		}
	}

	if win {
		RenderTemplate(w, "templates/Win.html", nil)
		return
	}

	// Structure des données envoyées au template
	data.HiddenWord = hiddenWord
	data.PhaseHangman = "/static/hangman-game-images/hangman-" + strconv.Itoa(6-lives) + ".svg"
	data.Lives = lives

	// Charger et afficher le template
	t, err := template.ParseFiles("templates/Play.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func CheckWin(hiddenWord string) bool {
	for _, char := range hiddenWord {
		if char == '_' {
			return false // Il reste des lettres à deviner
		}
	}
	return true // Le mot est complété
}
func RestartGame(w http.ResponseWriter, r *http.Request) {

}
