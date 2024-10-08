package hang

import (
	"log"
	"math/rand"
	"os"
	"strings"
)

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
		Lrand2 = rand.Intn(len(motJeu) - 1)
		for Lrand2 == Lrand1 {
			Lrand2 = rand.Intn(len(motJeu) - 1)
		}
	}
	if diff == 1 {
		Lrand2 = rand.Intn(len(motJeu) - 1)
		for Lrand2 == Lrand1 {
			Lrand2 = rand.Intn(len(motJeu) - 1)
		}
		Lrand3 = rand.Intn(len(motJeu) - 1)
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
