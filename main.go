package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var opts struct {
	NumRequests int `short:"r" long:"num-requests" description:"Number of requests to make" default:"1"`
	Concurrent  int `short:"c" long:"concurrent" description:"Number of concurrent connections to make" default:"1"`
}

func produceRequests(requestChan chan (string), target string, numRequests int) {
	for i := 0; i < numRequests; i++ {
		requestChan <- target
	}
	close(requestChan)
}

func consumeRequests(requestChan chan (string), resultChan chan (error)) {
	for target := range requestChan {
		doRequest(target, resultChan)
	}
}

func doRequest(target string, resultChan chan (error)) {
	response, err := http.Get(target)
	if err != nil {
		resultChan <- err
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		resultChan <- err
		return
	}

	resultChan <- nil
}

func main() {
	var err error
	var args []string

	parser := flags.NewParser(&opts, flags.Default)
	parser.Usage = "[options] URL"
	if args, err = parser.Parse(); err != nil {
		return
	}

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Requires one target URL")
		return
	}

	// Kick off workers
	target := args[0]
	requestChan := make(chan (string))
	resultChan := make(chan (error))

	for i := 0; i < opts.Concurrent; i++ {
		go consumeRequests(requestChan, resultChan)
	}

	startTime := time.Now()

	go produceRequests(requestChan, target, opts.NumRequests)

	// Gather results
	successCount := 0
	failureCount := 0

	for i := 0; i < opts.NumRequests; i++ {
		result := <-resultChan
		if result == nil {
			successCount++
		} else {
			failureCount++
		}
	}

	duration := time.Since(startTime)
	requestsPerSecond := float64(successCount) / duration.Seconds()

	// Output results
	fmt.Printf("# Success: %v\n", successCount)
	fmt.Printf("# Failure: %v\n", failureCount)
	fmt.Printf("%v requests/second\n", requestsPerSecond)
}
