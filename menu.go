package hang

import "github.com/gookit/color"

func ShowMenu() {
	color.Hex("#61e1f9").Println(`
			======================================
			1. Tapez 'L' pour deviner une lettre
			2. Tapez 'M' pour deviner un mot complet
			======================================
	`)
}
