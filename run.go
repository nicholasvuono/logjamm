package logjamm

import "sync"

//Options struct specifies the number of virtual users, test duration, or request iterations. In other words, the load tests cofiguration options.
type Options struct {
	Vus        int
	Duration   int
	Iterations int
}

func Run(options Options, requests []Request, test Test) ([]string, []map[string]int64) {
	var responses []string
	var timings []map[string]int64
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		responses = Batch(options, requests)
		wg.Done()
	}()
	go func() {
		timings = WebTest(options, test)
		wg.Done()
	}()
	wg.Wait()
	return responses, timings
}
