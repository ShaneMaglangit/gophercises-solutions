//go:generate stringer -type=Suit,Rank

package cards

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"
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

// Put can be used to add a card to the deck.
func (d *Deck) Put(c Card) {
	if exists := d.Contains(c); !exists {
		*d = append(*d, c)
	}
}

// Draw can be used to get the top card from the deck.
func (d *Deck) Draw() (Card, error) {
	// Check if the deck contains any card
	if len(*d) < 1 {
		return Card{}, errors.New("deck is empty")
	}

	// Return the last card in the deck
	c := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]
	return c, nil
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
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}

// New used to create a new deck of cards.
func New(opts ...func(Deck) Deck) Deck {
	deck := Deck{}

	// Create card for each suit
	for _, suit := range suits {
		// Create card for each rank
		for rank := minRank; rank <= maxRank; rank++ {
			deck.Put(Card{Suit: suit, Rank: rank})
		}
	}

	// Run functional options
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

// Jokers is used to add n jokers on the deck of card.
// Each joker card will have a rank based on their order of creation.
func Jokers(n int) func(Deck) Deck {
	return func(d Deck) Deck {
		for i := 0; i < n; i++ {
			d = append(d, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return d
	}
}

func Filter(f func(card Card) bool) func(Deck) Deck {
	return func(d Deck) Deck {
		var ret Deck
		for _, c := range d {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}
