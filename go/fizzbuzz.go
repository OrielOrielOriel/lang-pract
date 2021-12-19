package main

import (
    "fmt"
    "strings"
)

func main() {
    for i := 1; i <= 100; i++ {
        var write strings.Builder
        if i % 3 == 0 {
            write.WriteString("Fizz")
        }
        if i % 5 == 0 {
            write.WriteString("Buzz")
        }
        if write.String() != "" {
            fmt.Println(write.String())
        } else {
            fmt.Printf("%v\n", i)
        }
    }
}
