// declare a main package
package main

// import the fmt standard library, that contais many functions for formatting text
import "fmt"
import "rsc.io/quote"

func main() {
	// define the main funciont that executes by default when you run the main package
    fmt.Println(quote.Go())
}
