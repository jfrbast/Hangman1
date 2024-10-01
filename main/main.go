package main

import (
	"github.com/gookit/color"
	"hang"
	"time"
)

func main() {
	hang.Difficulty()
	hang.ChoixMot()
	color.Hex("#61e1f9").Println(" _   _                                                     \n| | | |   __ _   _ __     __ _   _ __ ___     __ _   _ __  \n| |_| |  / _` | | '_ \\   / _` | | '_ ` _ \\   / _` | | '_ \\ \n|  _  | | (_| | | | | | | (_| | | | | | | | | (_| | | | | |\n|_| |_|  \\__,_| |_| |_|  \\__, | |_| |_| |_|  \\__,_| |_| |_|\n                         |___/                             ")
	time.Sleep(500 * time.Millisecond)
	for {
		hang.Game()
	}
}
