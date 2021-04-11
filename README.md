# logjamm
[![Go Report Card](https://goreportcard.com/badge/github.com/logjamdev/logjam)](https://goreportcard.com/report/github.com/logjamdev/logjam)

E2E Load Testing Framework

!! CURRENTLY UNDER CONSTRUCTION !!

## Features

* Pure Go Library
* Simple and Readable API
* For use with Playwright (Go) - https://github.com/mxschmitt/playwright-go

## How it works

N/A

## Simplest Working Example

```go
package main

import (
	"fmt"

	"github.com/logjammdev/logjamm"
	"github.com/mxschmitt/playwright-go"
)

var options = logjamm.Options{
	Vus:        1,
	Duration:   10,
	Iterations: 0,
}

var requests = []logjamm.Request{
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

var test = func() map[string]int64 {

	timings := make(map[string]int64)

	pw, _ := playwright.Run()

	browser, _ := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})

	page, _ := browser.NewPage()

	label, duration, _ := logjamm.Step("Navigate to logjamm.io", func(page playwright.Page) playwright.Page {
		_, _ = page.Goto("https://logjamm.io")
		return page
	}, page)

	timings[label] = duration

	browser.Close()

	pw.Stop()

	return timings
}

func main() {
	responses, timings := logjamm.Run(options, requests, test)

	fmt.Println(responses)
	fmt.Println(timings)
}
```

## Simplest Working Example (Just Batch Request Functionality)

```go
package main

import (
	"fmt"

	"github.com/logjammdev/logjamm"
)

var options = logjamm.Options{
	Vus:        1,
	Duration:   0,
	Iterations: 5,
}

var requests = []logjamm.Request{
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

func main() {
	responses := logjamm.Batch(options, requests)

	fmt.Println(responses)
}
```

## Simplest Working Example (Just Web Test Functionality)

```go
package main

import (
	"fmt"

	"github.com/logjammdev/logjamm"
	"github.com/mxschmitt/playwright-go"
)

var options = logjamm.Options{
	Vus:        0,
	Duration:   0,
	Iterations: 5,
}

var test = func() map[string]int64 {

	timings := make(map[string]int64)

	pw, _ := playwright.Run()

	browser, _ := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})

	page, _ := browser.NewPage()

	label, duration, _ := logjamm.Step("Navigate to logjamm.io", func(page playwright.Page) playwright.Page {
		_, _ = page.Goto("https://logjamm.io")
		return page
	}, page)

	timings[label] = duration

	browser.Close()

	pw.Stop()

	return timings
}

func main() {
	responses, timings := logjamm.WebTest(options, test)

	fmt.Println(timings)
}
```


## To-Do

* ~~Implement WebDriver Functionality~~
* Add How it works details
* ~~Add Simplest Working Examples~~
* Add Request per Second to Options for Duration and Possibly Iteration (Batch Functionality)
* ~~Add Capture and Report End User Metrics/Timings API as Wrapper Function(WebDriver Functionality)~~
