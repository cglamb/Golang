package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// strcuture to hold the data
type ScrapedData struct {
	url         string   //website url
	Text        string   //text data
	MainHeaders string   //wiki headers
	SubHeaders  string   //wiki subheaders (if applcable)
	Bullets     string   // some wiki text is stored in bullets.  storing that text here
	Links       []string //grabs page links for future use in crawling
}

// cleans up scraped text to remove some wiki artifacts
func cleaner(text string) string {
	text = strings.ReplaceAll(text, "[edit]", " ") //removes the wiki function that allows user to edit
	return text
}

func scrapePage(url string, dataMap map[string]ScrapedData) {
	// iniate the colly instance
	c := colly.NewCollector()

	// setting a variable that uses the structure
	var data ScrapedData

	// html parsing instructions
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) { // On wiki pages, the body of text is in .mw-parser-output div
		data.url = url
		data.Text = cleaner(e.ChildText("p"))         // collects text
		data.MainHeaders = cleaner(e.ChildText("h2")) // collects main headers
		data.SubHeaders = cleaner(e.ChildText("h3"))  // collects sub headers
		data.Bullets = cleaner(e.ChildText("ul"))     //collects bullet data
		data.Links = e.ChildAttrs("a", "href")        //grabs links
	})

	// page to visit
	c.Visit(url)

	//putting the data into the map
	dataMap[url] = data
}

// marshalls and writes the json
func writeJson(dataMap map[string]ScrapedData, filename string) error {
	f, err := json.MarshalIndent(dataMap, "", "   ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, f, os.ModePerm)
	return err
}

// reads in a txt file containing all the urls we will scrape
func readUrl(filename string) ([]string, error) {
	txtContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	urls := strings.Split(string(txtContent), "\n")
	return urls, nil
}

func main() {

	dataMap := make(map[string]ScrapedData) //create a map to store data.  url will be the key

	// define variables to save user inputs
	var url_file string
	var f_name string

	// command line inputs for input and output files.
	flag.StringVar(&url_file, "input", "target_html.txt", "Path to the file containing URLs") //default to target_html.txt if no user input provided
	flag.StringVar(&f_name, "output", "output.jl", "Name of the output JSON file")            //default to output.json if no user input provided

	flag.Parse()

	// reading in a txt that contains urls to scarpe
	urls, err := readUrl(url_file)
	if err != nil {
		fmt.Println("Error reading in file:", err)
		return
	}

	// for every URL making an html request and parsing the html
	for _, url := range urls {
		scrapePage(url, dataMap)
	}

	// writes either the .jl
	err = writeJson(dataMap, f_name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON saved as:", f_name)

}
