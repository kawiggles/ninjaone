package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/kawiggles/ninjaone/internal/ninja"
)

type Model struct {
	client 	*ninja.Client
	menu 	Menu
	options []string
	offset	int
	height	int
	cursor	int
}

func NewModel(client *ninja.Client) Model {
	return Model{
		client: client,
		menu: MenuDevices,
		cursor: 0,
	}
}

type Menu int

const (
	MenuDevices Menu = iota
	MenuUsers
)

func (m Menu) GenString() string {
	switch m {
	case MenuDevices:
		return fmt.Sprintf("%-6s %-30s %-10s\n", "ID", "Name", "Status")
	case MenuUsers:
		return fmt.Sprintf("%-6s %-30s\n", "ID", "Name")
	}
	return "\n"
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

			if m.cursor < m.offset {
				m.offset--
			}

		case "down", "j":
			// TODO: Move cursor down
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

			if m.cursor >= m.offset+m.height {
				m.offset++
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
	devices, err := m.client.GetDevices()
	if err != nil {
		fmt.Printf("A thing went wrong when attempting to GetDevices(): %v\n", err)
	}

	for _, d := range devices {
		options = append(options, fmt.Sprintf("%-6d %-30s %-10s",
			d.ID, d.DisplayName, d.Maintenance.Status,))
	}

	var s strings.Builder
	s.WriteString(m.menu.GenString())

	for i, option := range m.visibleOptions() {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s.WriteString(fmt.Sprintf("%s %s\n", cursor, option))
	}

	s.WriteString("\nPress q to quit.\n")

	return tea.NewView(s.String())
}

func (m Model) visibleOptions() []string {
	end := m.offset + m.height

	if end > len(m.options) {
		end = len(m.options)
	}

	return m.options[m.offset:end]
}
