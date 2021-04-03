package adventure

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Arc struct {
	Title   string
	Story   []string
	Options []struct {
		Text string
		Arc  string
	}
}

func Adventure() {
	// Get the flags
	jsonFilename := flag.String("storyJson", "adventure/story.json", "contains the JSON data used to create the stories")
	flag.Parse()

	// Get the file
	file, err := os.Open(*jsonFilename)
	if err != nil {
		exit(fmt.Sprintln("An error encountered while trying to find the file ::", *jsonFilename))
		return
	}

	// Close the file later
	defer file.Close()

	// Read content
	jsonByte, err := ioutil.ReadAll(file)
	if err != nil {
		exit(fmt.Sprintln("An error encountered while trying to read the file", err))
		return
	}

	// Parse the Json content
	var arcs map[string]Arc
	err = json.Unmarshal(jsonByte, &arcs)
	if err != nil {
		exit(fmt.Sprintln("An error encountered while trying to parse the file ::", err))
		return
	}

	fmt.Println("Title", arcs["intro"].Title)

	for _, paragraph := range arcs["intro"].Story {
		fmt.Println(paragraph)
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
