package mtglib

import "log"

// Variable declaration
type Card struct {
	NAME string
	COST int
}

type Database struct {
	CARDS []Card{}
}

// FUNCTIONS
// function to open the database
//if there is none should create new files
func OpenDatabase() (Database, bool) {


	// Need to return a pointer(&)
}

// function to close the database
func (d *Database) CloseDatabase() bool {

}

// function to add a new card instance
func (d *Database) AddCard() bool {

}

// function to remove a card instance
func (d *Database) RemoveCard() bool {

}

// function that returns a card instance
func (d *Database) getCard() Card bool {

}

// function to change the count of a card
func (d *Database) ChangeCount(name string, number int) {

}