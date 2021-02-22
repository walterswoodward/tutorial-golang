package main

import (
    "fmt"
    "rsc.io/quote"
    "time"
)

func main() {
    fmt.Println("Initiating quote.go:")
    fmt.Println("The time is", time.Now())
    fmt.Println(quote.Hello())
    fmt.Println(quote.Go())
    fmt.Println(quote.Opt())
    fmt.Println(quote.Glass())
    fmt.Println("Goodbye!")
}