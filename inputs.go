package hang

import (
	"fmt"
	"github.com/gookit/color"
	"os"
)

func TryLetter() {
	var input string
	fmt.Println("Proposez une lettre : ")
	fmt.Scanln(&input)
	input = ToLower(input)

	if !IsAlpha(input) || len(input) != 1 {
		color.Hex("#ff0000").Println("Ceci n'est pas une lettre valide.")
		return
	}

	for _, l := range lettresEssayees {
		if l == input {
			fmt.Println("Vous avez déjà proposé cette lettre.")
			return
		}

	}
	lettresEssayees = append(lettresEssayees, input)

	found := false
	for i, char := range motJeu {
		if input == string(char) {
			motATrouver[i] = input
			found = true
		}
	}
	nbessais++
	if found {
		color.Hex("#2dc512").Println("Lettre trouvée !")

	} else {
		essaisRestants--
		color.Hex("#ff0000").Println("Lettre incorrecte. Essais restants :", essaisRestants)

	}

}

func TryMot() {
	var input string
	fmt.Println("Proposez le mot entier : ")
	fmt.Scanln(&input)
	input = ToLower(input)
	nbessais++
	if input == motJeu {
		fmt.Printf("Félicitations ! Vous avez deviné le mot :%v en %d essais \n", motJeu, nbessais)
		os.Exit(0)
	} else {
		essaisRestants -= 2
		color.Hex("#ff0000").Println("Mauvaise proposition. Essais restants :", essaisRestants)

	}
}

func Difficulty() {
	var input int

	fmt.Println("Choisissez le niveau de difficulté.")
	color.Hex("#2dc512").Println("1: Facile (10 essais + 3 indices)")
	color.Hex("#ffe400").Println("2: Moyen (7 essais + 2 indices)")
	color.Hex("#ff0000").Println("3: Difficile (5 essais + 1 indice)")
	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			color.Hex("#ff0000").Println("Erreur d'entrée, veuillez réessayer.")
			continue
		}
		switch input {
		case 1:
			color.Hex("#2dc512").Println("Vous avez choisi le niveau Facile.")
			essaisRestants = 10
			diff = 1
			return
		case 2:
			color.Hex("#ffe400").Println("Vous avez choisi le niveau Moyen.")
			essaisRestants = 8
			diff = 2
			return
		case 3:
			color.Hex("#ff0000").Println("Vous avez choisi le niveau Difficile.")
			essaisRestants = 6
			diff = 3
			return
		default:
			color.Hex("#ff0000").Println("Choix invalide, veuillez entrer 1, 2 ou 3.")
		}
	}
}
