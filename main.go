package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func affichage() {
	switch essaisRestants {
	case 10:
		fmt.Println(`
			
			

			
			
			
			_________
		`)
	case 9:
		fmt.Println(`
			
			

			|
			|
			|
			|_______
		`)
	case 8:
		fmt.Println(`
			
			_________
			|       
			|       
			|
			|
			|_______
		`)
	case 7:
		fmt.Println(`
			_________
			|       |
			|       
			|
			|
			|_______
		`)
	case 6:
		fmt.Println(`
			_________
			|       |
			|       O
			|
			|
			|_______
		`)
	case 5:
		fmt.Println(`
			_________
			|       |
			|       O
			|       |
			|
			|_______
		`)
	case 4:
		fmt.Println(`
			_________
			|       |
			|       O
			|      /|
			|
			|_______
		`)
	case 3:
		fmt.Println(`
			_________
			|       |
			|       O
			|      /|\
			|
			|_______
		`)
	case 2:
		fmt.Println(`
			_________
			|       |
			|       O
			|      /|\
			|      /
			|_______
		`)
	case 1:
		fmt.Println(`
			_________
			|       |
			|       O
			|      /|\
			|      / \
			|_______
		`)
	case 0:
		fmt.Println(`
			_________
			|       |
			|       O
			|      /|\
			|      / \
			|_______
			|   Perdu ! Le mot était : ` + motJeu)
	}
}

func ToLower(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		} else {
			result += string(char)
		}
	}
	return result
}

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'a' && s[i] <= 'z') {
			return false
		}
	}
	return true
}

// Variables globales
var tableauMot []string
var motATrouver []string
var motJeu string
var lettresEssayees []string
var essaisRestants = 10

func TryLetter() {
	var input string
	fmt.Println("Proposez une lettre : ")
	fmt.Scanln(&input)
	input = ToLower(input)

	if !IsAlpha(input) || len(input) != 1 {
		fmt.Println("Ceci n'est pas une lettre valide.")
		return
	}

	for _, l := range lettresEssayees {
		if l == input {
			fmt.Println("Vous avez déjà proposé cette lettre.")
			return
		}
		if input == "Retour" {
			fmt.Println("Retour.")
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

	if found {
		fmt.Println("Lettre trouvée !")
		affichage()
	} else {
		essaisRestants--
		fmt.Println("Lettre incorrecte. Essais restants :", essaisRestants)
		affichage()
	}

	fmt.Println("Mot à deviner : ", strings.Join(motATrouver, " "))
	fmt.Println("Lettres déjà proposées : ", strings.Join(lettresEssayees, ", "))
}

func TryMot() {
	var input string
	fmt.Println("Proposez le mot entier : ")
	fmt.Scanln(&input)
	input = ToLower(input)

	if input == motJeu {
		fmt.Println("Félicitations ! Vous avez deviné le mot :", motJeu)
		os.Exit(0)
	} else {
		essaisRestants -= 2
		fmt.Println("Mauvaise proposition. Essais restants :", essaisRestants)
		affichage()
	}
}

func ChoixMot() {
	data, err := os.ReadFile("Mots.txt")
	if err != nil {
		log.Fatal(err)
	}

	mots := strings.Split(strings.TrimSpace(string(data)), "\n")

	for i := range mots {
		mots[i] = strings.TrimSpace(mots[i])
	}
	tableauMot = mots

	motJeu = tableauMot[rand.Intn(len(tableauMot))]
	motATrouver = make([]string, len(motJeu))

	Lrand1 := rand.Intn(len(motJeu))
	Lrand2 := rand.Intn(len(motJeu))
	for Lrand2 == Lrand1 {
		Lrand2 = rand.Intn(len(motJeu))
	}

	for i, char := range motJeu {
		if i == Lrand1 || i == Lrand2 {
			motATrouver[i] = string(char)
		} else {
			motATrouver[i] = "_"
		}
	}
	fmt.Println("Mot à deviner : ", strings.Join(motATrouver, " "))
}

func main() {
	ChoixMot()
	fmt.Println(motJeu)

	for essaisRestants > 0 {
		var choix string
		fmt.Println("Tapez 'lettre'(ou L) pour proposer une lettre ou 'mot'(ou M) pour proposer un mot entier : ")
		fmt.Scanln(&choix)

		if choix == "lettre" || choix == "L" || choix == "l" {
			for {
				TryLetter()
			}

		} else if choix == "mot" || choix == "M" || choix == "m" {
			TryMot()
		} else {
			fmt.Println("Choix invalide. Réessayez.")
		}

		if strings.Join(motATrouver, "") == motJeu {
			fmt.Println("Félicitations ! Vous avez trouvé le mot :", motJeu)
			break
		}
	}

}
