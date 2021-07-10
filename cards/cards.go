//go:generate stringer -type=Suit,Rank

package cards

import (
	"errors"
	"fmt"
)

// Suit is used to define the given suit of the card,
// each suit is represented by an integer.
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

// array containing each usable suits
var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank is used to define the number within the given card from ace to king,
// each suit is represented by an integer.
type Rank uint8

const (
	// Skips 0
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// Used to iterate through the rank constants
const (
	minRank = Ace
	maxRank = King
)

// Card is used to define the a card within a deck
type Card struct {
	Suit
	Rank
}

// Returns the textual representation of the card
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// Deck is used as a collective struct for cards.
type Deck struct {
	Cards []Card
}

// Add can be used to add a card to the deck.
func (d *Deck) Add(c Card) {
	if _, exists := d.Contains(c); !exists {
		d.Cards = append(d.Cards, c)
	}
}

// Delete can be used to remove a card from the deck.
func (d *Deck) Delete(c Card) error {
	if i, exists := d.Contains(c); exists {
		d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
		return nil
	}
	return errors.New("card not found")
}

// Contains checks if a card exists within a given deck.
// Returns the index of the card in the list if it exists
// Returns -1 if card does not exists
func (d *Deck) Contains(c Card) (int, bool) {
	for i, card := range d.Cards {
		if c == card {
			return i, true
		}
	}
	return -1, false
}

// New used to create a new deck of cards
func New() Deck {
	var deck Deck

	// Create card for each suit
	for _, suit := range suits {
		// Create card for each rank
		for rank := minRank; rank <= maxRank; rank++ {
			deck.Add(Card{Suit: suit, Rank: rank})
		}
	}

	return deck
}
