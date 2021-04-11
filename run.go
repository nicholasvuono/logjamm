package logjam

//Options struct specifies the number of virtual users, test duration, or request iterations. In other words, the load tests cofiguration options.
type Options struct {
	Vus        int
	Duration   int
	Iterations int
}

func Run(options Options, requests []Request, test Test) {
	go Batch(options, requests)
	go WebTest(options, test)
}
