package hangman

var WordsList []string //Dictionnaire
var Word string        // Mot choisi au hasard

var word string // Mot à deviner
var hiddenWord string // Réinitialise le mot caché
var lives int

var data struct {
	HiddenWord string
	PhaseHangman string
	}
