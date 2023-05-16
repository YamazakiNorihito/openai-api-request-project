package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"

	"log"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    fmt.Println("Hello, World!")
	fmt.Println(quote.Go())


    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }
	
    fmt.Println(message)
}
