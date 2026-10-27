package tui

import (
	"github.com/kawiggles/ninjaone/internal/ninja"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	client 	*ninja.Client
}

func NewModel(client *ninja.Client) Model {
	return Model{
		client: client,
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

		case "down", "j":
			// TODO: Move cursor down

		case "left", "h":
			// TODO: Switch to left menu

		case "right", "l":
			// TODO: Switch to right menu

		case "enter":
			_, ok := m.selected[m.cursor]
		}
	}
}

func (m Model) View() tea.View {
	return tea.NewView("hi")
}
