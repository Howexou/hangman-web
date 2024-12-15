package hangman

type Jeu struct {
	Word string
	Found string
	NbDeVie int
	DejaMis []rune
}

var LeJeu Jeu