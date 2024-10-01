package hang

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"strings"
	"time"
)

func Game() {
	for essaisRestants > 0 {
		ShowMenu()
		Affichage()
		time.Sleep(1 * time.Second)
		fmt.Println("Lettres déjà proposées : ", strings.Join(lettresEssayees, ", "))
		time.Sleep(1 * time.Second)
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
			fmt.Println("  ____                     __   _ \n / ___| __ _  __ _ _ __   /_/  | |\n| |  _ / _` |/ _` | '_ \\ / _ \\ | |\n| |_| | (_| | (_| | | | |  __/ |_|\n \\____|\\__,_|\\__, |_| |_|\\___| (_)\n             |___/                ")
			fmt.Printf("Félicitations ! Vous avez trouvé le mot : %v en %d essais", motJeu, nbessais)
			os.Exit(0)
		}
	}

	if essaisRestants == 0 {
		fmt.Println(" ____              _         _ \n|  _ \\ ___ _ __ __| |_   _  | |\n| |_) / _ \\ '__/ _` | | | | | |\n|  __/  __/ | | (_| | |_| | |_|\n|_|   \\___|_|  \\__,_|\\__,_| (_)")
		fmt.Println("Désolé, vous avez perdu. Le mot était :", motJeu)
		os.Exit(0)
	}
}
