package main

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func affichage() {
	time.Sleep(1 * time.Second)
	switch essaisRestants {
	case 10:
		color.Hex("#808080").Println(`

			_________

		`)
		time.Sleep(1 * time.Second)
	case 9:
		color.Hex("#e295ff").Println(`

			|
			|
			|
			|_______
		`)
		time.Sleep(1 * time.Second)
	case 8:
		color.Hex("#808080").Println(`

			_________
			|       
			|       
			|
			|
			|_______
		`)
		time.Sleep(1 * time.Second)
	case 7:
		color.Hex("#808080").Println(`

			_________
			|       |
			|       
			|
			|
			|_______
		`)
		time.Sleep(1 * time.Second)
	case 6:
		color.Hex("#e295ff").Println(`

			_________
			|       |
			|       O
			|
			|
			|_______
		`)
		time.Sleep(1 * time.Second)
	case 5:
		color.Hex("#808080").Println(`

			_________
			|       |
			|       O
			|       |
			|
			|_______
		`)
		time.Sleep(1 * time.Second)
	case 4:
		color.Hex("#808080").Println(`

			_________
			|       |
			|       O
			|      /|
			|
			|_______
		`)
		time.Sleep(1 * time.Second)
	case 3:
		color.Hex("#808080").Println(`

			_________
			|       |
			|       O
			|      /|\
			|
			|_______
		`)
		time.Sleep(1 * time.Second)
	case 2:
		color.Hex("#808080").Println(`

			_________
			|       |
			|       O
			|      /|\
			|      /
			|_______
		`)
		time.Sleep(1 * time.Second)
	case 1:
		color.Hex("#808080").Println(`

			_________
			|       |
			|       O
			|      /|\
			|      / \
			|_______
		`)
		time.Sleep(1 * time.Second)
	case 0:
		color.Hex("#808080").Println(`

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

func difficulty() {
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

// Variables globales
var diff int
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

	if found {
		color.Hex("#2dc512").Println("Lettre trouvée !")
		affichage()
	} else {
		essaisRestants--
		color.Hex("#ff0000").Println("Lettre incorrecte. Essais restants :", essaisRestants)
		affichage()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Mot à deviner : ", strings.Join(motATrouver, " "))
	time.Sleep(1 * time.Second)
	fmt.Println("Lettres déjà proposées : ", strings.Join(lettresEssayees, ", "))
	time.Sleep(1 * time.Second)
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
		color.Hex("#ff0000").Println("Mauvaise proposition. Essais restants :", essaisRestants)
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

	var Lrand1 int
	var Lrand2 int
	var Lrand3 int

	if diff == 2 {
		Lrand2 = rand.Intn(len(motJeu))
		for Lrand2 == Lrand1 {
			Lrand2 = rand.Intn(len(motJeu))
		}
	}
	if diff == 1 {
		Lrand2 = rand.Intn(len(motJeu))
		for Lrand2 == Lrand1 {
			Lrand2 = rand.Intn(len(motJeu))
		}
		Lrand3 = rand.Intn(len(motJeu))
		for Lrand3 == Lrand1 && Lrand3 == Lrand2 {
			Lrand3 = rand.Intn(len(motJeu))
		}
	}

	for i, char := range motJeu {
		if i == Lrand1 || i == Lrand2 || i == Lrand3 {
			motATrouver[i] = string(char)
		} else {
			motATrouver[i] = "_"
		}
	}
}

func showMenu() {
	color.Hex("#61e1f9").Println(`
			======================================
			 _    _                                          
			| |  | |                                         
		 	| |__| | __ _ _ __   __ _ _ __ ___   __ _ _ __    
			|  __  |/ _` + "`" + ` | '_ \ / _` + "`" + ` | '_ ` + "`" + ` _ \ / _` + "`" + ` | '_ \   
		 	| |  | | (_| | | | | (_| | | | | | | (_| | | | |  
		 	|_|  |_|\__,_|_| |_|\__, |_| |_| |_|\__,_|_| |_|  
					    __/  |
					    |___/                  
			======================================
			1. Tapez 'L' pour deviner une lettre
			2. Tapez 'M' pour deviner un mot complet
			======================================
	`)
}

func main() {
	difficulty()
	ChoixMot()
	time.Sleep(2 * time.Second)
	for essaisRestants > 0 {
		showMenu()
		color.Hex("#1694d8").Println("Mot à deviner : ", strings.Join(motATrouver, " "))
		time.Sleep(1 * time.Second)
		var choix string
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choix)

		if choix == "lettre" || choix == "L" || choix == "l" {
			TryLetter()
		} else if choix == "mot" || choix == "M" || choix == "m" {
			TryMot()

		} else {
			color.Hex("ff0000").Println("Choix invalide. Réessayez.")
		}

		if strings.Join(motATrouver, "") == motJeu {
			fmt.Println("Félicitations ! Vous avez trouvé le mot :", motJeu)
			break
		}
	}

	if essaisRestants == 0 {
		fmt.Println("Désolé, vous avez perdu. Le mot était :", motJeu)
	}
}
