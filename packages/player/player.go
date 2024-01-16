package player

import "fmt"

type Player struct {
	Pseudo string
	Pawn   string
}

type Game struct {
	Players []Player
}

func GetPlayer() {
	var (
		player      Player
		nbPlayer    int  = 2
		playerOne   bool = true
		colorPlayer int
	)

	game := Game{
		Players: make([]Player, 0),
	}

	for i := 0; i < nbPlayer; i++ {
		fmt.Print("Entrer votre pseudo <= à 10 charactère: ")
		_, err := fmt.Scan(&player.Pseudo)
		if err != nil {
			fmt.Println("Errur lors de la saisie: ", err)
			return
		}

		if len(player.Pseudo) > 10 {
			fmt.Println("Veuillez entrer une pseudo <= à 10 charactère !")
			continue
		}

		if playerOne {
			fmt.Println("Choisissez la couleur du pion (1 pour 🔵, 2 pour 🔴) : ")
			_, err = fmt.Scan(&colorPlayer)
			if err != nil {
				fmt.Println("Errur lors de la saisie: ", err)
				return
			}

			if colorPlayer != 1 && colorPlayer != 2 {
				fmt.Println("Choisissez un chiffre entre 1 et 2")
			}

			if colorPlayer == 1 {
				player.Pawn = "🔵"
			} else {
				player.Pawn = "🔴"
			}

			game.Players = append(game.Players, Player{player.Pseudo, player.Pawn})

			playerOne = false

		} else {
			if colorPlayer == 1 {
				player.Pawn = "🔴"
			} else {
				player.Pawn = "🔵"
			}
			game.Players = append(game.Players, Player{player.Pseudo, player.Pawn})
		}
	}

	// Affichage des joueurs dans le jeu
	for _, player := range game.Players {
		fmt.Printf("Joueur : %s, Pion : %s\n", player.Pseudo, player.Pawn)
	}

}
