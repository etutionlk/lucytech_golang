package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type Site struct {
	URL               string
	HTMLVersion       string
	PageTitle         string
	HeadingCount      map[string]int
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks []string
	ContainsLoginForm bool
}

func (s *Site) getHtmlContent() string {
	response, err := http.Get(s.URL)
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return " "
	}
	return string(body)

}

func (s *Site) getHtmlVersion(siteContent string) {

	contentUpper := strings.ToUpper(siteContent)

	start := strings.Index(contentUpper, "<!DOCTYPE")
	end := strings.Index(contentUpper[start:], ">")
	doctype := contentUpper[start : start+end+1]

	switch {
	case doctype == "<!DOCTYPE HTML>":
		s.HTMLVersion = "HTML5"
	case strings.Contains(doctype, "HTML 4.01"):
		s.HTMLVersion = "HTML 4.01"
	case strings.Contains(doctype, "XHTML 1.0"):
		s.HTMLVersion = "XHTML 1.0"
	case strings.Contains(doctype, "XHTML 1.1"):
		s.HTMLVersion = "XHTML 1.1"
	default:
		s.HTMLVersion = "Unknown or Custom DOCTYPE"
	}
}

func (s *Site) getPageTitle(siteContent string) {
	doc, err := html.Parse(strings.NewReader(siteContent))
	if err != nil {
		fmt.Println("Error parsing HTML")
	}
	s.PageTitle = traverse(doc)
}

func traverse(n *html.Node) string {
	var title string

	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		title = n.FirstChild.Data
		return title
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = traverse(c)
		if title != "" {
			return title
		}
	}

	return ""
}

func (s *Site) getHeadingCount(siteContent string) {
	doc, err := html.Parse(strings.NewReader(siteContent))
	if err != nil {
		fmt.Println("Error parsing HTML")
	}

	headingCount := make(map[string]int)

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "h1", "h2", "h3", "h4", "h5", "h6":
				headingCount[n.Data]++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)

	tageCount := make(map[string]int)
	for tag, count := range headingCount {
		tageCount[tag] = count
	}

	s.HeadingCount = tageCount

}

func (s *Site) getAllLinks(siteContent string) []string {
	doc, err := html.Parse(strings.NewReader(siteContent))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return []string{""}
	}

	var links []string

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)

	return links

}

func (s *Site) getInaccessibleLinks(links []string) {
	var badLinks []string

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	for _, link := range links {
		resp, err := client.Head(link)
		if err != nil || resp.StatusCode >= 400 {
			badLinks = append(badLinks, link)
			continue
		}
		resp.Body.Close()
	}

	s.InaccessibleLinks = badLinks
}

func (s *Site) isContainsLoginForm() {
	// code
}

func main() {
	// url := "https://go.dev/doc/tutorial/getting-started"

	url := "https://go.dev/doc/tutorial/getting-started#install"

	s := Site{URL: url}

	// fmt.Println(s.getHtmlContent())

	content := s.getHtmlContent()

	// s.getHtmlVersion(content)

	// fmt.Println(s.HTMLVersion)

	s.getPageTitle(content)
	s.getHeadingCount(content)

	links := s.getAllLinks(content)
	// fmt.Println(links)

	s.getInaccessibleLinks(links)

	fmt.Println(s)

}
