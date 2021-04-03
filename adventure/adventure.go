package adventure

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
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

var arcs map[string]Arc

func Adventure() {
	// Get the flags
	jsonFilename := flag.String("storyJson", "adventure/story.json", "contains the JSON data used to create the stories")
	flag.Parse()

	// Extract the stories
	err := loadStories(jsonFilename)
	if err != nil {
		exit("", err)
	}

	// Start the server
	http.HandleFunc("/", defaultHandler)
	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// Get the slug from the url
	slug := r.URL.Path[len("/"):]

	// Redirect to intro as the default page
	if slug == "" {
		http.Redirect(w, r, "intro", http.StatusFound)
		return
	}

	// Parse the HTML template
	pageTemplate, _ := template.ParseFiles("adventure/story.html")

	// Load the page
	if content, ok := arcs[slug]; ok {
		_ = pageTemplate.Execute(w, content)
	}
}

func loadStories(jsonFilename *string) error {
	// Get the file
	file, err := os.Open(*jsonFilename)
	if err != nil {
		return err
	}

	// Close the file later
	defer file.Close()

	// Read content
	jsonByte, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// Parse the Json content
	err = json.Unmarshal(jsonByte, &arcs)
	if err != nil {
		return err
	}

	return nil
}

func exit(msg string, err error) {
	fmt.Println(msg, err)
	os.Exit(1)
}
