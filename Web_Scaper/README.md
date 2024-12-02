# Go-Based Web Scraping using the Colly Framework

- Author: Charles Lamb
- Contact Info: charlamb@gmail.com
- GitHub address for this project: [https://github.com/cglamb/Web_Scraper](https://github.com/cglamb/Web_Scraper)
- Git Clone command for the repository: `git clone https://github.com/cglamb/Web_Scraper.git`

## Introduction

This project develops an application that scrapes text information from Wikipedia URLs. The application is written in Go(lang) and relies on the Colly Framework. URLs to be scraped are passed to the application via a `.txt` file (see `target_html.txt`). An HTML request to that URL is made via Colly, and data is parsed using Colly's `OnHTML` function. The end result is a JSON (`.jl`) file output.

## Structure of the Output Data

The JSON key is the URL addresses of each site being visited. The JSON also contains values for the following data elements:

- `text`: This is the text from the body of the wiki page.
- `main_headers`: These are the main header names within the wiki page.
- `sub_headers`: Sub-headers under main headers.
- `bullets`: Text contained in bullets on the page. Some wiki pages have bullets, which the HTML code stores differently from the main body text.  This captures that text.
- `links`: These are links referenced in the body of the wiki page.  These might be valuable for future using in web crawling.

## Running the Command Line Application

If the terminal's current directory is the directory containing the executable, the program can be run from the command line using the following command: `./web-scraper -input target_html.txt -output output.jl`. The `-output` can be changed to the desired location and filename for the output text. The default location is the current directory, and the default name is `output.jl`. The `-input` can be changed to the location and name of the CSV file being read. The default location is the current directory, and the default file name is `target_html.txt`.

## Explanation of Files

- `output.jl`: JSON output containing scraped information.
- `scraper.go`: Codebase
- `scaper_test.go`: Testing for functions in `scraper.go`.
- `target_html.txt`: TXT containing URLs meant to be read in.
- `web-scraper.ext`: Executable web scraping program.
- `go.sum`: File used to manage dependencies.
- `go.mod`: Module definition file for managing system dependencies.
- `logs/`
  - `benchmark_log.txt`: Log file for benchmark testing.
  - `executable_log.txt`: Log showing running of the command-line application in the terminal.
  - `test_log.txt`: Log showing running of test functions.

## Compiling Instructions

The executable can be compiled in Go via the `go build` command within the terminal.

## Testing

Test routines were developed for two functions that read URLs from the text file and do some light data scrubbing of parsed data. The author manually reviewed the application for making HTML requests and parsing HTML data, but no test functions were developed in Go.

## Enhancements

Given the program is making requests and parsing multiple URLs, goroutines could be added to allow for scraping multiple URLs concurrently.
