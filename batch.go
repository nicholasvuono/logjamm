package logjam

import (
	"errors"
	"fmt"
	"sync"
	"time"

	httpbatch "github.com/logjamdev/http-batch"
)

//Request struct mimicking the one found in the wreckhttp library
type Request = httpbatch.Request

//Global wait group for go routines within this test
var wg sync.WaitGroup

//Global string slice to capture all responses
var responses []string

//Batch is a function to send batch requests based on specified load test configuration options
func Batch(options Options, requests []Request) []string {
	if options.Iterations != 0 && options.Duration == 0 {
		concurrrentBatchIterations(options, requests)
	} else if options.Iterations == 0 {
		concurrentBatchDuration(options, requests)
	} else {
		err := errors.New("error Options: Duration and Iteration cannot be used at the same time")
		explain(err)
	}
	return responses
}

//Sends batch requests concurently based on the duration specified
func concurrentBatchDuration(options Options, requests []Request) {
	after := time.Now().Add(time.Duration(options.Duration) * time.Second)
	for {
		now := time.Now()
		for i := 0; i < options.Vus; i++ {
			wg.Add(1)
			go sendBatch(requests)
		}
		if now.After(after) {
			break
		}
		wg.Wait()
	}
}

//Sends batch requests concurrently based on the number of iterations specified
func concurrrentBatchIterations(options Options, requests []Request) {
	for i := 0; i < options.Iterations; i++ {
		for i := 0; i < options.Vus; i++ {
			wg.Add(1)
			go sendBatch(requests)
		}
		wg.Wait()
	}
}

//Logic to actually send the batch requests and append responses to the global slice
func sendBatch(requests []Request) {
	defer wg.Done()
	batch, err := httpbatch.Batch(requests)
	explain(err)
	responses = append(responses, fmt.Sprintf("%v", batch.Send()))
}
