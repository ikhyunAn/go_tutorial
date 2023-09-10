package main

import (
    "fmt"

    "example/greetings"
)

func main() {
    // Get greetings message from module: greetings
    message := greetings.Hello("Ikhyun")
    fmt.Println("Hello, World!")
    fmt.Println(message)
}
