package cards

import (
	"reflect"
	"testing"
)

func TestCard_String(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want string
	}{
		{"Ace Suit", Card{Heart, Ace}, "Ace of Hearts"},
		{"Number Suit", Card{Spade, Three}, "Three of Spades"},
		{"Face Suit", Card{Diamond, King}, "King of Diamonds"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_Add(t *testing.T) {
	tests := []struct {
		name string
		deck Deck
		card Card
		want []Card
	}{
		{"Empty deck", Deck{[]Card{}}, Card{Heart, Ace}, []Card{{Heart, Ace}}},
		{"Add card", Deck{[]Card{{Heart, Ace}}}, Card{Spade, Three}, []Card{{Heart, Ace}, {Spade, Three}}},
		{"No duplicate", Deck{[]Card{{Heart, Ace}}}, Card{Heart, Ace}, []Card{{Heart, Ace}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deck.Add(tt.card)
			if !reflect.DeepEqual(tt.deck.Cards, tt.want) {
				t.Errorf("Add() got = %v, want %v", tt.deck.Cards, tt.want)
			}
		})
	}
}

func TestDeck_Delete(t *testing.T) {
	tests := []struct {
		name     string
		deck     Deck
		card     Card
		want     []Card
		hasError bool
	}{
		{"Empty deck", Deck{[]Card{}}, Card{Heart, Ace}, []Card{}, true},
		{"Remove existing card", Deck{[]Card{{Heart, Ace}}}, Card{Heart, Ace}, []Card{}, false},
		{"Remove non-existing card", Deck{[]Card{{Heart, Ace}}}, Card{Heart, Jack}, []Card{{Heart, Ace}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.deck.Delete(tt.card)
			if !reflect.DeepEqual(tt.deck.Cards, tt.want) {
				t.Errorf("Delete() got = %v, want %v", tt.deck.Cards, tt.want)
			}
			if tt.hasError == (err == nil) {
				t.Error("Delete() expected an error")
			}
			if tt.hasError != (err != nil) {
				t.Errorf("Delete() unexpected error: %v", err)
			}
		})
	}
}

func TestDeck_Contains(t *testing.T) {
	tests := []struct {
		name   string
		deck   Deck
		card   Card
		want   int
		exists bool
	}{
		{"Empty deck", Deck{[]Card{}}, Card{Diamond, Ace}, -1, false},
		{"Card exists", Deck{[]Card{{Diamond, Ace}}}, Card{Diamond, Ace}, 0, true},
		{"Card does not exists", Deck{[]Card{{Diamond, King}}}, Card{Diamond, Ace}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, exists := tt.deck.Contains(tt.card)
			if i != tt.want {
				t.Errorf("Contains() index = %v, want %v", i, tt.want)
			}
			if exists != tt.exists {
				t.Errorf("Contains() exists = %v, want %v", exists, tt.exists)
			}
		})
	}
}

func TestNew(t *testing.T) {
	deck := New()
	// 13 ranks * 4 suits
	if n := len(deck.Cards); n != 13*4 {
		t.Errorf("%d, want %d", n, 13*4)
	}
}
