package hang

import (
	"github.com/gookit/color"
	"time"
)

func Affichage() {
	time.Sleep(1 * time.Second)
	switch essaisRestants {
	case 10:
		color.Hex("#808080").Println(`

			_________

		`)
		time.Sleep(1 * time.Second)
	case 9:
		color.Hex("#808080").Println(`

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
		color.Hex("#808080").Println(`

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
