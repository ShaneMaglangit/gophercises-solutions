//go:generate stringer -type=Suit,Rank

package cards

import (
	"fmt"
	"sort"
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

// Used to iterate through the rank constants.
const (
	minRank = Ace
	maxRank = King
)

// Card is used to define the a card within a deck.
type Card struct {
	Suit
	Rank
}

// Returns the textual representation of the card.
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank, c.Suit)
}

// Deck is used as a collective struct for cards.
type Deck []Card

// Add can be used to add a card to the deck.
func (d *Deck) Add(c Card) {
	if exists := d.Contains(c); !exists {
		*d = append(*d, c)
	}
}

// Delete can be used to remove a card from the deck.
func (d *Deck) Delete(c Card) bool {
	if i := d.Find(c); i != -1 {
		*d = append((*d)[:i], (*d)[i+1:]...)
		return true
	}
	return false
}

// Find gets the index of the card in the deck. Returns -1 if it doesn't exist.
func (d Deck) Find(c Card) int {
	for i, card := range d {
		if c == card {
			return i
		}
	}
	return -1
}

// Contains checks if a card exists within a given deck.
func (d Deck) Contains(c Card) bool {
	for _, card := range d {
		if c == card {
			return true
		}
	}
	return false
}

// Shuffle sorts the cards in the deck in random order.
func (d *Deck) Shuffle() {

}

// New used to create a new deck of cards.
func New(opts ...func(Deck) Deck) Deck {
	var deck Deck

	// Create card for each suit
	for _, suit := range suits {
		// Create card for each rank
		for rank := minRank; rank <= maxRank; rank++ {
			deck.Add(Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		deck = opt(deck)
	}

	return deck
}

// DefaultSort sorts the card in a deck based on their absolute rank value
// calculated by suit * maxRank + rank
func DefaultSort(d Deck) Deck {
	sort.Slice(d, Less(d))
	return d
}

// Sort is for more generalized sorting dictated by the less function parameter.
func Sort(less func(d Deck) func(i, j int) bool) func(Deck) {
	return func(d Deck) {
		sort.Slice(d, Less(d))
	}
}

// Less creates a less function that will be used for sorting.
// It is used to check whether if card A is lower than card B.
func Less(d Deck) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(d[i]) < absRank(d[j])
	}
}

// absRank is used for returning the absolute rank of a given card.
// This will be used as the comparison function for the cards.
func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
