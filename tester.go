package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	var input string
  fmt.Println(input)
	for {
    fmt.Scan()
    if input != "" {
      break
    }
  }
	fmt.Println(input)
}
