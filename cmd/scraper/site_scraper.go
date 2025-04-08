package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Site struct {
	URL               string
	HTMLVersion       string
	PageTitle         string
	HeadingCount      map[string]int
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks int
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
	c := traverse(doc)
	fmt.Println(c)
}

func traverse(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}

}

func (s *Site) getHeadingCount() {
	// code
}

func (s *Site) getAllLinks() {
	// code
}

func (s *Site) getInaccessibleLinks() {
	// code
}

func (s *Site) isContainsLoginForm() {
	// code
}

func main() {
	url := "https://example.com"

	s := Site{URL: url}

	// fmt.Println(s.getHtmlContent())

	content := s.getHtmlContent()

	// s.getHtmlVersion(content)

	// fmt.Println(s.HTMLVersion)

	s.getPageTitle(content)

}
