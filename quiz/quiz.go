package quiz

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Quiz(answers []string) int {
	// Open the file
	file, err := os.Open("C:\\Users\\shane\\Workspace\\gophercises\\quiz\\problems.csv")
	if err != nil {
		fmt.Println("An error has been encountered ::", err)
		return -1
	}

	// Read the file
	reader := csv.NewReader(file)
	questions, _ := reader.ReadAll()

	// Initialize variables
	corr := 0

	// Iterate through each question
	for i, question := range questions[:10] {
		// Show the question
		fmt.Printf("%v: ", question[0])

		// Get the input
		var answer string

		if answers == nil {
			_, _ = fmt.Scan(&answer)
		} else {
			answer = answers[i]
		}

		// Check if answer is correct
		if answer == question[1] {
			corr++
		}
	}

	// Display the score
	fmt.Printf("Correct Answers: %v, Wrong Answers: %v\n", corr, 10-corr)
	_ = file.Close()

	return corr
}
