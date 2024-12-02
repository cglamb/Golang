package main

import (
	"os"
	"reflect"
	"testing"
)

func TestCleaner(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"This is a [edit] test", "This is a   test"},      //[edit] is removed
		{"[edit] [edit] [edit]", "     "},                  // multiple [edit] are removed
		{"No [edit] here", "No   here"},                    // [edit] in the middle of the text
		{"No [Edit] here", "No [Edit] here"},               // Case-sensitive test
		{"This is a [ed]it test", "This is a [ed]it test"}, // [ed]it is not removed
	}

	for _, test := range tests {
		result := cleaner(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}

// did not build a test case on scrapePage
// some example of setting up test serves avaialble at
// Saha, Amit. 2022. Practical Go: Building Scalable Network & Non-network Applications. New York: Wiley.
// instead i tested on live wiki pages and ensured the results returned were expected

func TestReadUrl(t *testing.T) {

	// create sample import file
	urlsContent := `https://en.wikipedia.org/wiki/Robotics
https://en.wikipedia.org/wiki/Robot
https://en.wikipedia.org/wiki/Reinforcement_learning
https://en.wikipedia.org/wiki/Robot_Operating_System
https://en.wikipedia.org/wiki/Intelligent_agent
https://en.wikipedia.org/wiki/Software_agent
https://en.wikipedia.org/wiki/Robotic_process_automation
https://en.wikipedia.org/wiki/Chatbot
https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence
https://en.wikipedia.org/wiki/Android_(robot)`

	//write the temp file
	err := os.WriteFile("test_urls.txt", []byte(urlsContent), os.ModePerm)
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove("test_urls.txt") // remove the test file

	// call our function to read test_urls
	urls, err := readUrl("test_urls.txt")
	if err != nil {
		t.Fatalf("Error reading URLs from file: %v", err)
	}

	// Define the expected URLs
	expectedURLs := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	// Check if the result matches the expected URLs
	if !reflect.DeepEqual(urls, expectedURLs) {
		t.Errorf("Unexpected URLs. Expected: %+v, Got: %+v", expectedURLs, urls)
	}
}

func BenchmarkScrapePage(b *testing.B) {
	dataMap := make(map[string]ScrapedData)

	url := "https://en.wikipedia.org/wiki/Robotics"

	for i := 0; i < b.N; i++ {
		scrapePage(url, dataMap)
	}
}
