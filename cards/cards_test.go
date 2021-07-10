package cards

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	deck := New()
	// 13 ranks * 4 suits
	if n := len(deck.Cards); n != 13*4 {
		t.Errorf("%d, want %d", n, 13*4)
	}
}

func TestCard_String(t *testing.T) {
	type fields struct {
		Suit Suit
		Rank Rank
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Ace Suit", fields{Heart, Ace}, "Ace of Hearts"},
		{"Number Suit", fields{Spade, Three}, "Three of Spades"},
		{"Face Suit", fields{Diamond, King}, "King of Diamonds"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				Suit: tt.fields.Suit,
				Rank: tt.fields.Rank,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeck_Add(t *testing.T) {
	type fields struct {
		Cards []Card
	}
	type args struct {
		c Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Card
	}{
		{"Empty deck", fields{[]Card{}}, args{Card{Heart, Ace}}, []Card{{Heart, Ace}}},
		{"Add card", fields{[]Card{{Heart, Ace}}}, args{Card{Spade, Three}}, []Card{{Heart, Ace}, {Spade, Three}}},
		{"No duplicate", fields{[]Card{{Heart, Ace}}}, args{Card{Heart, Ace}}, []Card{{Heart, Ace}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deck{
				Cards: tt.fields.Cards,
			}
			d.Add(tt.args.c)
			if !reflect.DeepEqual(d.Cards, tt.want) {
				t.Errorf("Add() got = %v, want %v", d.Cards, tt.want)
			}
		})
	}
}

func TestDeck_Delete(t *testing.T) {
	type fields struct {
		Cards []Card
	}
	type args struct {
		c Card
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     []Card
		hasError bool
	}{
		{"Empty deck", fields{[]Card{}}, args{Card{Heart, Ace}}, []Card{}, true},
		{"Remove existing card", fields{[]Card{{Heart, Ace}}}, args{Card{Heart, Ace}}, []Card{}, false},
		{"Remove non-existing card", fields{[]Card{{Heart, Ace}}}, args{Card{Heart, Jack}}, []Card{{Heart, Ace}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deck{
				Cards: tt.fields.Cards,
			}
			err := d.Delete(tt.args.c)
			if !reflect.DeepEqual(d.Cards, tt.want) {
				t.Errorf("Delete() got = %v, want %v", d.Cards, tt.want)
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
	type fields struct {
		Cards []Card
	}
	type args struct {
		c Card
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
		want1  bool
	}{
		{"Empty deck", fields{[]Card{}}, args{Card{Diamond, Ace}}, -1, false},
		{"Card exists", fields{[]Card{{Diamond, Ace}}}, args{Card{Diamond, Ace}}, 0, true},
		{"Card does not exists", fields{[]Card{{Diamond, King}}}, args{Card{Diamond, Ace}}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Deck{
				Cards: tt.fields.Cards,
			}
			got, got1 := d.Contains(tt.args.c)
			if got != tt.want {
				t.Errorf("Contains() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Contains() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
