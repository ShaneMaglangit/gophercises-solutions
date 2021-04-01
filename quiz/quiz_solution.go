package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func QuizSol() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	score := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == problem.answer {
			fmt.Println("Correct!")
			score++
		}
	}

	fmt.Printf("Total score %d out of %d", score, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
