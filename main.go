package main

import (
	"flag"
	"log"
	"regexp"
	"strings"
)

var (
	url string
)

func main() {
	flag.StringVar(&url, "url", "", "The URL to analyse")
	flag.Parse()

	// FOR DEBUGGING PURPOSES!
	url = "https://www.haevg-rz.de/"

	if url == "" {
		log.Fatal("No URL specified!")
	}

	// prepare for use in regex
	preparedUrl := prepareUrl(url)

	regex, err := regexp.Compile(preparedUrl)
	if err != nil {
		log.Fatalf("could not compile regex '%s': '%s'", preparedUrl, err.Error())
	}

	checkWebsite(url, regex)

}

func checkWebsite(u string, regex *regexp.Regexp) {
	// jede URL nur 1x verwendbar!
	content, status, err := downloadWebsiteContent(u)
	if err != nil {
		log.Fatalf("could not download website ")
	}

	if status >= 400 {
		return
	}

	matches := regex.FindAllString(content, -1)
	if len(matches) == 0 {
		return
	}

	for _, m := range matches {
		checkWebsite(m, regex)
	}
}

func prepareUrl(u string) string {
	// colon as well?
	u = strings.ReplaceAll(u, `/`, `\/`)
	u = strings.ReplaceAll(u, `.`, `\.`)
	return u
}

func downloadWebsiteContent(url string) (string, int, error) {

}
