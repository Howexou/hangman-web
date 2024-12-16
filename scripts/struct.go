package hangman

type Jeu struct {
	Word string
	Found string
	NbDeVie int
	DejaMis []rune
}

var LeJeu Jeu

/* Ici on a la structure du jeu
Word est le mot a deviner
Found est l'état des lettres à découvrir
NbDeVie Bah c'est le nombre de vie
DejaMis la liste de lettre tenté par le joueur*/