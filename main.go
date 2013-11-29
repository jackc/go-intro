package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"net/http"
	"os"
	"time"
)

var opts struct {
	NumRequests int `short:"r" long:"num-requests" description:"Number of requests to make" default:"1"`
}

func doRequest(target string) error {
	response, err := http.Get(target)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		return err
	}

	return nil
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

	for i := 0; i < opts.NumRequests; i++ {
		go doRequest(target)
	}

	time.Sleep(time.Second)
}
