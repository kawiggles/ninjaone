package tui

import (
	"fmt"

	"github.com/kawiggles/ninjaone/internal/ninja"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	client 	*ninja.Client
	options []string
	cursor	int
}

func NewModel(client *ninja.Client) Model {
	return Model{
		client: client,
		cursor: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			// TODO: Move cursor up
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			// TODO: Move cursor down
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case "left", "h":
			// TODO: Switch to left menu

		case "right", "l":
			// TODO: Switch to right menu

		case "enter":
			// TODO: Enter selected menu
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	s := "Select a list\n\n"

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, option)
	}

	s += "\nPress q to quit.\n"

	return tea.NewView("hi")
}
