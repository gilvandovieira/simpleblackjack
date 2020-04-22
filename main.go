package main

import (
	"fmt"

	"github.com/gilvandovieira/simpleblackjack/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
