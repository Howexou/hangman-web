package hangman

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	LeJeu = Jeu{}
	RenderTemplate(w, "Home", nil)
}

func HandleRequests() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", Hangman) 
	http.ListenAndServe(":8081", nil)
}

func Hangman(w http.ResponseWriter, r *http.Request) {
	if LeJeu.Word == "" {
		InitGame()
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		letter := rune(r.FormValue("letter")[0])
		ProcessLetter(letter)
	}

	data := struct {
		Found    string
		NbDeVie  int
		DejaMis  []rune
		Word     string
	}{
		Found:    LeJeu.Found,
		NbDeVie:  LeJeu.NbDeVie,
		DejaMis:  LeJeu.DejaMis,
		Word:     LeJeu.Word,
	}

	RenderTemplate(w, "Hangman", data)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template : "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
