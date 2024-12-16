package hangman

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	LeJeu = Jeu{}
	RenderTemplate(w, "Home", nil)
}
/* (w http.ResponseWriter, pour écrire la réponse http qui sera envoyer au web
r *http.Request) c'est la requête reçu par le web*/

func HandleRequests() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", Hangman) 
	http.ListenAndServe(":8081", nil)
}
/* la func HandleResquest sert à démarrer le serveur HTTP

http.HandleFunc("/", Home)
http.HandleFunc("/hangman", Hangman)

http.handlefunc sert a associé une url à une fonction
les url associé sont Home et Hangman du coup
*/

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

/* la function sert a initialiser le jeu  

if r.Method == http.MethodPost {
	r.ParseForm()
	letter := rune(r.FormValue("letter")[0])
	ProcessLetter(letter)
}

r.methode == , c'est pour vérifier la requête est une requête POST, ce type de requête est utiliser pour envoyer une lettre proposer
r.paseform analyse le formulaire POST envoyé avec la requête
r.formvalue récupere la valeur du champ letter et retourne une châine de caractére
ProcessLetter sert à mettre à jour l'état du jeu

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

on a déjà expliquer dans struct.go à quoi sa servait, c'est les data pour le rendu du template HTML

RenderTemplate(w, "Hangman", data)

charge le fichier template/hangman et injecte les données 
*/

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template : "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
 /*la function renderTemplate charge un fichier HTML depuis le répertoire
 
 func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{})

 w hhtp.responseWriter permet d'nevoyer des  réponse HTML au web
 tmlp string le nom du fichier template
 data interface les données qui seront injecté dans le template

 t, err := template.ParseFiles("templates/" + tmpl + ".html")

 Pour charger le fichier de template
 template.Parsefile analyse le fichier HTML spécifié, ici il charge un seul fichier 
 
 if err != nil 

 gestion des erreur

 t.Execute(w, data)

 exécute le template chaargé avec les données fournis dans le paramétre data
 */