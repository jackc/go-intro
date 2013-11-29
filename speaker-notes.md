= Intro

Why learn a programming language?

* To use it
* To learn something and take it back to another language

= s1 - Hello, World!

* package
* import / namespace
* func main
* case determines visibility

= s2 - Command line arguments

* Println can use reflection to print non-string arguments
* Slices - slice of an arrays
* Demonstrate auto-format

= s3 - curl

* Initialization assignment - type inferance
* Multiple returns
* Verbose error handling
* defer
* Demonstrate checked slice index access by running without providing a command line argument
* io.Copy - discuss reader and writer interfaces
* Remember to start sinatra hello world app

= s4 - Multiple requests

* External packages
* var keyword
* struct tags
* single use struct instead of defining a type
* for loop - only loop construct

= s5 - Extract doRequest

* Note repetitive if err != nil
* Demonstrate no warnings / only errors
* Demonstrate IDE like auto-complete

= s6 - Concurrent requests

* Go routines
** Light-weight "threads" multiplexed on real threads
** Return value is lost
** No communication back from go routines makes sleep necessary

= s7 - Channels

* Channels are the concurrent safe means for go routines to communicate

= s8 - More channels

* Use channels to provide input and give output
* Control how many simultaneous requests in progress at once

= s9 - Add timing

* Add sleep to sinatra to demonstrate that concurrency is working

= Go's Use It Reasons

Go is designed to be used, not as an academic exercise.

* Concurrency primitives built into language (channels and go routines)
* Static type checking / strict error handling
* Statically linked binary output with minimal dependencies
* Cross-platform
* Excellent standard library
* Fast enough that it can eliminate or defer the need for caching, one the the famously hard problems
* Minimal memory usage
* Compiles are fast
* Simple
* Static duck typing

= Go's Learn Something Reasons

* Pointers
* Possibly channels, but they would be hard to use elsewhere without language support

= Negatives for Go

* No REPL
* Verbose error handling / mixed blessing
* Static typing makes certain tasks like testing and freeform hash munging more verbose
* No templates/generics

