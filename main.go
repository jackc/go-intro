package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var opts struct {
	NumRequests int `short:"r" long:"num-requests" description:"Number of requests to make" default:"1"`
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

	target := args[0]

	resultChan := make(chan (error))

	for i := 0; i < opts.NumRequests; i++ {
		go doRequest(target, resultChan)
	}

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

	fmt.Printf("# Success: %v\n", successCount)
	fmt.Printf("# Failure: %v\n", failureCount)
}
