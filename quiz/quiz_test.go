package quiz

import "testing"

func TestQuizAllCorrect(t *testing.T) {
	answers := []string{"10", "2", "11", "3", "14", "4", "5", "6", "5", "6"}
	score := Quiz(answers)
	if score != 10 {
		t.Fatalf(`Score expected: 10, Received: %v`, score)
	}
}

func TestQuizNoCorrect(t *testing.T) {
	answers := []string{"1", "3", "2", "4", "5", "6", "7", "8", "9", "10"}
	score := Quiz(answers)
	if score != 0 {
		t.Fatalf(`Score expected: 0, Received: %v`, score)
	}
}

func TestQuizPartialCorrect(t *testing.T) {
	answers := []string{"10", "2", "11", "3", "14", "6", "7", "8", "9", "10"}
	score := Quiz(answers)
	if score != 5 {
		t.Fatalf(`Score expected: 5, Received: %v`, score)
	}
}
