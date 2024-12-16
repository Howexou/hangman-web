package hangman

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Home", nil)
}

// Configure les routes HTTP
func HandleRequests() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/play", Play)
	http.HandleFunc("/win", Win)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
	http.ListenAndServe(":8081", nil)
}

// Afficher le HTML avec des données "dynamiques"
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) { 
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

// Gère les requêtes HTTP pour la route "/"
func handleIndex(w http.ResponseWriter, r *http.Request) { 
	randomWord := PickWord()
	fmt.Println("Mot aléatoire généré : ", randomWord)

	WebData := WebData{
		Word: Word,
	}

	RenderTemplate(w, Template, WebData)

}

func Win(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "victory", nil)
}
