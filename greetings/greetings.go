package greetings

import "fmt"

//Hello should return a greeting from the other person
func Hello(name string) string {
    //Returns a greeting
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
