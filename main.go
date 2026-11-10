package main

import (
	"fmt"
	"os"

	"github.com/kawiggles/ninjaone/internal/ninja"
	"github.com/kawiggles/ninjaone/internal/tui"
	tea "charm.land/bubbletea/v2"
)

func main() {
	client := ninja.NewClient()

	program := tea.NewProgram(tui.NewModel(client))
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
		os.Exit(1)
	}
}
