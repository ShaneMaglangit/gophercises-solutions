package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func Quiz() {
	csvFilename := flag.String("csv", "quiz/problems.csv", "a csv file in the format of 'question, answer")
	timeLimit := flag.Int("timeLimit", 30, "duration of the quiz in seconds")
	flag.Parse()

	// Open the file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// Read the file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	// Initialize variables
	problems := parseLines(lines)
	score := 0

	// Create timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//go func() {
	//	<-timer.C
	//	endQuiz(score, len(problems))
	//}()

	for i, problem := range problems {
		// Display the question
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)

		// Create answer channel
		answerCh := make(chan string)

		// Goroutine function that accepts the answer from user
		go func() {
			var answer string
			_, _ = fmt.Scanf("%s", &answer)

			// Pass answer to channel
			answerCh <- answer
		}()

		// Do something if we get something from the channel
		select {
		case <-timer.C:
			fmt.Printf("\nTotal score %d out of %d", score, len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.answer {
				score++
			}
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
