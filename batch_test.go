package logjamm

import (
	"fmt"
	"strconv"
	"testing"
)

var tests = []func(t *testing.T){
	TestBatch,
}

var options = Options{
	Vus:        3,
	Duration:   0,
	Iterations: 10,
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

func TestAll(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i), test)
	}
}
