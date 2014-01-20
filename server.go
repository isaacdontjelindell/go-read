package main

import (
    "fmt"

    "bufio"

    "os"
)

func main() {
    
    fmt.Print("Enter a title to search for: ")
    bio := bufio.NewReader(os.Stdin)
    line, _, _ := bio.ReadLine()

    s := string(line[:])

    fmt.Println(s)
}
