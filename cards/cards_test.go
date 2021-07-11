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

func TestDeck_Put(t *testing.T) {
	tests := []struct {
		name string
		deck Deck
		card Card
		want Deck
	}{
		{"Empty deck", Deck{}, Card{Heart, Ace}, Deck{{Heart, Ace}}},
		{"Put card", Deck{{Heart, Ace}}, Card{Spade, Three}, Deck{{Heart, Ace}, {Spade, Three}}},
		{"No duplicate", Deck{{Heart, Ace}}, Card{Heart, Ace}, Deck{{Heart, Ace}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.deck.Put(tt.card)
			if !reflect.DeepEqual(tt.deck, tt.want) {
				t.Errorf("Put() got = %v, want %v", tt.deck, tt.want)
			}
		})
	}
}

func TestDeck_Draw(t *testing.T) {
	tests := []struct {
		name     string
		deck     Deck
		want     Card
		hasError bool
	}{
		{"Fresh Deck", New(), Card{Heart, King}, false},
		{"2 Cards", Deck{{Heart, King}, {Diamond, Ace}}, Card{Diamond, Ace}, false},
		{"Empty Deck", Deck{}, Card{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.deck.Draw()
			if got != tt.want {
				t.Errorf("Draw() got = %v, want %v", got, tt.want)
			}
			if tt.hasError && err == nil {
				t.Errorf("Draw() expected an error")
			}
			if !tt.hasError && err != nil {
				t.Errorf("Draw() unexpected error: %v", err)
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
			if exists := tt.deck.Contains(tt.card); exists != tt.want {
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
			if s := DefaultSort(tt.d); s[0] != tt.cFirst {
				t.Errorf("DefaultSort() first card = %v, want %v", s[0], tt.cFirst)
			}
		})
	}
}

func TestDeck_Shuffle(t *testing.T) {
	tests := []struct {
		name string
		d    Deck
	}{
		{"Fresh Deck", New()},
		{"Fresh Deck", Deck{{Heart, Ace}, {Spade, Three}, {Diamond, King}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dCopy := make(Deck, len(tt.d))
			copy(dCopy, tt.d)
			dCopy.Shuffle()
			if reflect.DeepEqual(dCopy, tt.d) {
				t.Errorf("Shuffle() order remains unchanged")
			}
		})
	}
}

func TestJokers(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"0 Jokers", 0, 0},
		{"3 Jokers", 3, 3},
		{"-1 Jokers", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New(Jokers(tt.n))
			n := 0
			for _, c := range d {
				if c.Suit == Joker {
					n++
				}
			}
			if n != tt.want {
				t.Errorf("Jokers() = %v, want %v", n, tt.want)
			}
		})
	}
}
