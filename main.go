package main

import (
    "fmt"
    build "./builder"
)

func main() {
   fmt.Println("test")
   b := build.SelectStmt{}
   fmt.Println(b)
}
