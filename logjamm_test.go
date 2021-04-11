package logjamm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/mxschmitt/playwright-go"
)

var tests = []func(t *testing.T){
	TestBatch,
}

var options = Options{
	Vus:        5,
	Duration:   0,
	Iterations: 2,
}

var requests = []Request{
	{
		Method:  "GET",
		URL:     "https://httpbin.org/get",
		Headers: nil,
		Body:    nil,
	},
	{
		Method: "POST",
		URL:    "https://httpbin.org/post",
		Headers: map[string][]string{
			"Accept": {"application/json"},
		},
		Body: map[string]string{
			"name":  "Test API Guy",
			"email": "testapiguy@email.com",
		},
	},
}

func TestBatch(t *testing.T) {
	responses := Batch(options, requests)
	if len(responses) == 0 {
		t.Error("Testing Error: number of responses is zero!")
	}
	fmt.Println(responses)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(len(responses))
}

var test = func() map[string]int64 {

	timings := make(map[string]int64)

	pw, _ := playwright.Run()

	browser, _ := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})

	page, _ := browser.NewPage()

	label, duration, _ := Step("Navigate to logjamm.io", func(page playwright.Page) playwright.Page {
		_, _ = page.Goto("https://logjamm.io")
		return page
	}, page)

	timings[label] = duration

	browser.Close()

	pw.Stop()

	return timings

}

func TestWebTest(t *testing.T) {
	timings := WebTest(options, test)
	if timings[0]["Navigate to logjamm.io"] == 0 {
		t.Error("Testing Error: Navigate to logjamm.io step does not have a duration associated")
	}
	fmt.Println(timings)
}

func TestRun(t *testing.T) {
	responses, timings := Run(options, requests, test)
	if responses == nil || timings == nil {
		t.Error("Testing Error: nil values detected")
	}
	fmt.Println(responses)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(timings)
}

func TestAll(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i), test)
	}
}
