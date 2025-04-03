package main

import (
	"fmt"
	"net/http"
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

func (s *Site) getHtmlContent() {
	// code
	resp, err := http.Get(s.URL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("I am here", resp)
}

func (s *Site) getHtmlVersion() {
	// code goes here
}

func (s *Site) getPageTitle() {
	// code
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
	url := "https://examplensdss.com"

	s := Site{URL: url}

	s.getHtmlContent()
}
