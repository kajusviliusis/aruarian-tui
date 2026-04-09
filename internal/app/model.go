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
	switch typed := msg.(type) {
	case menu.SelectionMsg:
		switch typed.Selection {
		case menu.SelectionNotes:
			// todo: launch nvim
			return m, nil
		case menu.SelectionTodo:
			m.state = TodoState
			return m, nil
		case menu.SelectionTimer:
			m.state = TimerState
			return m, nil
		}
	}

	switch m.state {
	case MenuState:
		var cmd tea.Cmd
		m.menu, cmd = m.menu.Update(msg)
		return m, cmd
	case TodoState, TimerState:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			m.state = MenuState
		}
		return m, nil
	default:
		return m, nil
	}
}

func (m Model) View() string {
	switch m.state {
	case MenuState:
		return m.menu.View()
	case TodoState:
		return "TODO\n\nplaceholder\n\nesc: back to menu\n"
	case TimerState:
		return "TIMER\n\nplaceholder\n\nesc: back to menu\n"
	default:
		return "aruarian-tui\n\nunknown state\n"
	}
}
