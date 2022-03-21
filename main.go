package main

import (
    "fmt"
    "os"
)

func main() {
    source := os.Args[1]
    wordsCount := 0
    for _, ch := range []rune(source) {
        if wordsCount == 0 {
            wordsCount++
        }
        if ch == ' ' {
            wordsCount++
        }
    }
    fmt.Println(wordsCount)
}