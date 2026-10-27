package main

import (
	"fmt"

	"github.com/kawiggles/ninjaone/internal/ninja"
	"github.com/kawiggles/ninjaone/internal/tui"
)

func main() {
	client := ninja.NewClient()

	model := tui.NewModel(client)
}
