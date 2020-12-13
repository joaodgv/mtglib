package mtglib

import "log"

type Card struct {
	NAME string
	COST int
}

func PrintCard(c Card) {
	log.Println(c)
}