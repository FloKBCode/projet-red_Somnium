package main

import (
	"fmt"
	"somnium/game"
)

var gradient = []string{
	"\033[38;2;70;100;150m",   // Bleu clair pour visibilité sur fond sombre
	"\033[38;2;92;111;145m",   // Bleu acier
	"\033[38;2;123;44;191m",   // Violet brumeux
	"\033[38;2;157;78;221m",   // Violet vif
	"\033[38;2;192;132;252m",  // Violet clair
}

var asciiArt = []string{
	" ███████╗  ██████╗  ███╗   ███╗ ███╗   ██╗ ██╗  ██╗   ██╗ ███╗   ███╗",
	" ██╔════╝ ██╔═══██╗ ████╗ ████║ ████╗  ██║ ██║  ██║   ██║ ████╗ ████║",
	" ███████╗ ██║   ██║ ██╔████╔██║ ██╔██╗ ██║ ██║  ██║   ██║ ██╔████╔██║",
	" ╚════██║ ██║   ██║ ██║╚██╔╝██║ ██║╚██╗██║ ██║  ██║   ██║ ██║╚██╔╝██║",
	" ███████║ ╚██████╔╝ ██║ ╚═╝ ██║ ██║ ╚████║ ██║  ╚██████╔╝ ██║ ╚═╝ ██║",
	" ╚══════╝  ╚═════╝  ╚═╝     ╚═╝ ╚═╝  ╚═══╝ ╚═╝   ╚═════╝  ╚═╝     ╚═╝",
}

func main() {
	lineCount := len(asciiArt)
	gradCount := len(gradient)

	for i, line := range asciiArt {
		// Dégradé vertical : ligne haute = couleur foncée, ligne basse = couleur claire
		colorIndex := i * (gradCount - 1) / (lineCount - 1)
		color := gradient[colorIndex]

		fmt.Println(color + line + "\033[0m") // Appliquer couleur à toute la ligne
	}

	game.MainMenu()
}
