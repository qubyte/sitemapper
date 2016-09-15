package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/purell"
	"github.com/qubyte/sitemap"
)

func checkFlags() (*url.URL, int) {
	startURL := flag.String("start-url", "", "A fully qualified URL for the crawl to begin with.")
	jobs := flag.Int("jobs", 1, "Number of simultaneous requests to allow.")
	flag.Parse()

	if *startURL == "" {
		log.Fatalln("A start-url to crawl is required.")
	}

	if *jobs < 1 {
		log.Fatal("The job count cannot be less than 1.")
	}

	normalizedOriginString, err := purell.NormalizeURLString(*startURL, purell.FlagsSafe)

	if err != nil {
		log.Fatal("Unable to normalize start-url.")
	}

	normalizedURL, err := url.Parse(normalizedOriginString)

	if err != nil || !normalizedURL.IsAbs() {
		log.Fatal("start-url must be a fully qualified URL.")
	}

	return normalizedURL, *jobs
}

func main() {
	originURL, jobs := checkFlags()

	log.Println("Domain:", originURL.String())
	log.Println("Workers:", jobs)

	done := make(chan bool)

	sitemap := sitemap.NewSiteMap(originURL)

	go sitemap.Crawl(jobs, done)

	<-done

	result, _ := json.Marshal(sitemap)

	fmt.Println(string(result))
}
