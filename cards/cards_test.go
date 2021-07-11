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
		want Deck
	}{
		{"Empty deck", Deck{}, Card{Heart, Ace}, Deck{{Heart, Ace}}},
		{"Add card", Deck{{Heart, Ace}}, Card{Spade, Three}, Deck{{Heart, Ace}, {Spade, Three}}},
		{"No duplicate", Deck{{Heart, Ace}}, Card{Heart, Ace}, Deck{{Heart, Ace}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deck.Add(tt.card)
			if !reflect.DeepEqual(tt.deck, tt.want) {
				t.Errorf("Add() got = %v, want %v", tt.deck, tt.want)
			}
		})
	}
}

func TestDeck_Delete(t *testing.T) {
	tests := []struct {
		name    string
		deck    Deck
		card    Card
		want    Deck
		success bool
	}{
		{"Empty deck", Deck{}, Card{Heart, Ace}, Deck{}, false},
		{"Remove existing card", Deck{{Heart, Ace}}, Card{Heart, Ace}, Deck{}, true},
		{"Remove non-existing card", Deck{{Heart, Ace}}, Card{Heart, Jack}, Deck{{Heart, Ace}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			success := tt.deck.Delete(tt.card)
			if !reflect.DeepEqual(tt.deck, tt.want) {
				t.Errorf("Delete() got = %v, want %v", tt.deck, tt.want)
			}
			if success != tt.success {
				t.Errorf("Delete() success = %v, want %v", success, tt.success)
			}
		})
	}
}

func TestDeck_Contains(t *testing.T) {
	tests := []struct {
		name string
		deck Deck
		card Card
		want bool
	}{
		{"Empty deck", Deck{}, Card{Diamond, Ace}, false},
		{"Card exists", Deck{{Diamond, Ace}}, Card{Diamond, Ace}, true},
		{"Card does not exists", Deck{{Diamond, King}}, Card{Diamond, Ace}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists := tt.deck.Contains(tt.card)
			if exists != tt.want {
				t.Errorf("Contains() exists = %v, want %v", exists, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	deck := New()
	// 13 ranks * 4 suits
	if n := len(deck); n != 13*4 {
		t.Errorf("New() got = %d, want %d", n, 13*4)
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name   string
		d      Deck
		cFirst Card
	}{
		{"New() Sort", New(), Card{Spade, Ace}},
		{"Two Cards", Deck{Card{Heart, Jack}, Card{Spade, Ace}}, Card{Spade, Ace}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DefaultSort(tt.d)
			if s[0] != tt.cFirst {
				t.Errorf("DefaultSort() first card = %v, want %v", s[0], tt.cFirst)
			}
		})
	}
}
