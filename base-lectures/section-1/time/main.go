package main

import (
	"fmt"
	"time"
)

func main() {
	second := time.Second
	fmt.Print(int64(second / time.Millisecond))
}
