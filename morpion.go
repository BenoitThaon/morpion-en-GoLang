package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// déclaration du damier
var damier = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

// déclaration validation du coup
var coupOk bool

//déclaration entrée utilisateur
var inUser string

// Joueur
var player int

// Marque du joueur ( x ou o)
var marque string

// Jeu gagné
var win bool

func main() {

	// création scanner
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			player = 1
			marque = "x"
		} else {
			player = 2
			marque = "o"
		}
		// remise à zero de la validation du coup
		coupOk = false

		displaySpace()

		for !coupOk {
			displayDamier(damier)

			fmt.Print("Joueur ", player, " entrez un nombre compris entre 1 à 9 :")
			scanner.Scan()

			// récupération de l'entrée utilisateur
			inUser = scanner.Text()

			// valideCoup(inUser, damier[:])

			coupOk, damier = valideCoup(inUser, damier, marque)

		}
		if isWin(damier, marque) {
			displayDamier(damier)
			fmt.Println("Le joueur ", player, " à gagné !")
			win = true
			break
		}
	}

	if !win {
		displayDamier(damier)
		fmt.Println("Partie nulle !")
	}

}

// fonction affiche la grille
func displayDamier(damier [9]string) {
	fmt.Println(damier[:3])
	fmt.Println(damier[3:6])
	fmt.Println(damier[6:])
}

// fonction espace
func displaySpace() {
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
}

// fonction vérifie si entrée utilisateur est bonne (compris entre 1 et 9 ; pas deja joué donc encore dispo dans le tableau)
func valideCoup(inUser string, damier [9]string, marque string) (bool, [9]string) {

	inUserInt, err := strconv.Atoi(inUser)
	if inUserInt > 0 && inUserInt < 10 {

		if damier[inUserInt-1] == inUser {
			damier[inUserInt-1] = marque
			return true, damier
		} else {
			fmt.Println("Cette case est déjà prise !")
		}

	} else if err == nil {
		fmt.Println("Votre nombre doit être compris entre 0 à 9 !")
	} else {
		fmt.Println("Entrez un nombre et non autre chose !")
	}

	return false, damier
}

// fonction vérifie si le jeu est gagné
func isWin(damier [9]string, marque string) bool {

	switch {
	case damier[0] == marque && damier[1] == marque && damier[2] == marque:
		return true
	case damier[3] == marque && damier[4] == marque && damier[5] == marque:
		return true
	case damier[6] == marque && damier[7] == marque && damier[8] == marque:
		return true
	case damier[0] == marque && damier[3] == marque && damier[6] == marque:
		return true
	case damier[1] == marque && damier[4] == marque && damier[7] == marque:
		return true
	case damier[2] == marque && damier[5] == marque && damier[8] == marque:
		return true
	case damier[0] == marque && damier[4] == marque && damier[8] == marque:
		return true
	case damier[6] == marque && damier[4] == marque && damier[2] == marque:
		return true
	default:
		return false
	}
}
