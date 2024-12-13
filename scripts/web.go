package hangman

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "Home", nil)
}

func HandleRequests() {
	http.HandleFunc("/", handleIndex)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))
	http.ListenAndServe(":8081", nil)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(Template)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	randomWord := PickWord()
	fmt.Println("Mot aléatoire généré : ", randomWord)

	WebData := WebData{
		Word: Word,
	}

	RenderTemplate(w, Template, WebData)

}
