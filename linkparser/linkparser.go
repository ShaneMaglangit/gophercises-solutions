package linkparser

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"os"
	"strings"
)

type Link struct {
	href string
	text string
}

func LinkParser() {
	// Get the flags
	htmlFilename := flag.String("html", "linkparser/sample.html", "HTML file containing the source code for the page")
	flag.Parse()

	file, err := ioutil.ReadFile(*htmlFilename)
	if err != nil {
		exit(fmt.Sprintln("An error has been encountered while loading file ::", err))
		return
	}

	node, err := html.Parse(strings.NewReader(string(file)))
	if err != nil {
		exit(fmt.Sprintln("An error has been encountered while loading file ::", err))
		return
	}

	links := extractLink(node)
	fmt.Println(links)
}

func extractLink(n *html.Node) []Link {
	if n == nil {
		return nil
	}

	var links []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, Link{a.Val, extractText(n.FirstChild)})
			}
		}
	}

	links = append(links, extractLink(n.FirstChild)...)
	links = append(links, extractLink(n.NextSibling)...)

	return links
}

func extractText(n *html.Node) string {
	if n == nil {
		return ""
	}

	var nodeText string
	if n.Type == html.TextNode {
		nodeText = strings.TrimSpace(n.Data)
	}

	return fmt.Sprint(nodeText, extractText(n.FirstChild), extractText(n.NextSibling))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
