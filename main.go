package main

import (
	"fmt"
)

func main() {
	client := NewClient()
	fmt.Print(client.token)
}
