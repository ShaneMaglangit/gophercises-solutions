package sitemap

import (
	"encoding/xml"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Sitemap() {
	// Get the flags
	siteUrl := flag.String("siteUrl", "https://calhoun.io/", "Site to generate the sitemap for")
	flag.Parse()

	links := []string{*siteUrl}
	for i := 0; i < len(links); i++ {
		pageSource, err := loadPage(links[i])

		if err != nil {
			log.Fatal(fmt.Sprintln("An error has been encountered ::", err))
			return
		}

		foundLinks := extractLink(pageSource, getDomain(*siteUrl))

		for _, link := range foundLinks {
			if !checkIfExists(links, link) {
				links = append(links, link)
			}
		}
	}

	createXML(links)
}

func checkIfExists(slice []string, newString string) bool {
	for _, element := range slice {
		if element == newString {
			return true
		}
	}
	return false
}

func loadPage(url string) (*html.Node, error) {
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	source, err := html.Parse(strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}

	return source, nil
}

func getDomain(rawUrl string) string {
	u, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatal(fmt.Sprintln("An error has been encountered ::", err))
		return ""
	}
	return u.Host
}

func extractLink(n *html.Node, defaultDomain string) []string {
	if n == nil {
		return nil
	}

	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				switch {
				case string(a.Val[0]) == "/":
					links = append(links, fmt.Sprintf("https://%s%s", defaultDomain, a.Val))
				case getDomain(a.Val) == defaultDomain:
					links = append(links, a.Val)
				}
			}
		}
	}

	links = append(links, extractLink(n.FirstChild, defaultDomain)...)
	links = append(links, extractLink(n.NextSibling, defaultDomain)...)

	return links
}

func createXML(links []string) {
	type url struct {
		Loc string `xml:"loc"`
	}

	// Map links to XML format
	var urlset []url
	for _, link := range links {
		urlset = append(urlset, url{link})
	}

	urlXML, _ := xml.MarshalIndent(urlset, "", "\t")
	prefix := xml.Header + `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`
	suffix := "</urlset>"
	sitemapXML := []byte(prefix + string(urlXML) + suffix)

	err := os.WriteFile("sitemap/sitemap.xml", sitemapXML, 0666)
	if err != nil {
		fmt.Println(err)
	}
}
