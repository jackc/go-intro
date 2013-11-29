package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"net/http"
	"os"
)

var opts struct {
	NumRequests int `short:"r" long:"num-requests" description:"Number of requests to make" default:"1"`
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
		response, err := http.Get(target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		defer response.Body.Close()

		_, err = io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}
