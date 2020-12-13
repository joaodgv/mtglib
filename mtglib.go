package mtglib

import "log"

type struct Card {
	NAME string
	COST int
}

func PrintCard(c Card) {
	log.Println(c)
}