package hangman

var WordsList []string //Dictionnaire
var Word string        // Mot choisi au hasard

var word string // Mot à deviner
var hiddenWord string // Réinitialise le mot caché
var lives int // Actualise le nombre de vies 

var data struct {
	HiddenWord string
	PhaseHangman string
	Lives int
}
