package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/kajusviliusis/aruarian-tui/internal/menu"
)

type Model struct {
	state AppState
	menu  menu.Model
}

func NewModel() Model {
	return Model{
		state: MenuState,
		menu: menu.NewModel([]string{
			"NOTES",
			"TODO",
			"TIMER",
			"QUIT",
		}),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case MenuState:
		var cmd tea.Cmd
		m.menu, cmd = m.menu.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m Model) View() string {
	switch m.state {
	case MenuState:
		return m.menu.View()
	default:
		return "aruarian-tui\n\nunknown state\n"
	}
}
