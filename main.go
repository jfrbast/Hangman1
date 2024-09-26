package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
		if !(s[i] >= 97 && s[i] <= 122) {
			return false
		}
	}
	return true
}

//

//

var tableauMot []string
var motatrouver []string
var MotJeu = tableauMot[0]
var Tablettre []string
var Compteur = 0

//

//

func ChoixMot() {

	data, err := os.ReadFile("Mots.txt")
	if err != nil {
		log.Fatal(err)
	}
	mot := string(data)
	tableauMot = strings.SplitAfter(mot, "\n")

	var Random int
	Random = rand.Intn(len(tableauMot))
	MotJeu = tableauMot[Random]
	fmt.Println(MotJeu)

	Lrand1 := rand.Intn(len(MotJeu))
	Lrand2 := rand.Intn(len(MotJeu))
	if Lrand2 == Lrand1 {
		if Lrand2 == len(MotJeu) {
			Lrand2 -= 1
		} else {
			Lrand2 += 1
		}
	}
	for i, char := range MotJeu {
		if i == Lrand1 || i == Lrand2 {
			motatrouver = append(motatrouver, string(char))
		} else {
			motatrouver = append(motatrouver, string("_"))
		}

	}
	fmt.Println(motatrouver)
}

func main() {
	ChoixMot()
	var Input string
	fmt.Scanf(Input)
	ToLower(Input)
	for {
		if IsAlpha(Input) == false {
			fmt.Println("c'est pas une lettre.")
			time.Sleep(2 * time.Second)
			continue
		} else {
			for _, char := range MotJeu {
				if Input == string(char) {
					fmt.Println("Lettre trouvée")
					time.Sleep(1 * time.Second)
					motatrouver[char] = Input
				} else {
					fmt.Println("Pas la bonne lettre")
					Tablettre = append(Tablettre, string(Input))
					time.Sleep(1 * time.Second)
				}
			}

		}
	}

}
