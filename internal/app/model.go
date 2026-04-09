package app

import (
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kajusviliusis/aruarian-tui/internal/menu"
	"github.com/kajusviliusis/aruarian-tui/internal/timer"
	"github.com/kajusviliusis/aruarian-tui/internal/todo"
)

type Model struct {
	state AppState
	menu  menu.Model
	timer timer.Model
	todo  todo.Model
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
		timer: timer.NewModel(25 * time.Minute),
		todo:  todo.NewModel(todoStorePath()),
	}
}

func todoStorePath() string {
	cfgDir, err := os.UserConfigDir()
	if err != nil || cfgDir == "" {
		return "todos.json"
	}

	appDir := filepath.Join(cfgDir, "aruarian-tui")
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		return "todos.json"
	}

	return filepath.Join(appDir, "todos.json")
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
	case TodoState:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			m.state = MenuState
			return m, nil
		}

		var cmd tea.Cmd
		m.todo, cmd = m.todo.Update(msg)
		return m, cmd
	case TimerState:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" {
			m.timer = m.timer.Pause()
			m.state = MenuState
			return m, nil
		}

		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m Model) View() string {
	switch m.state {
	case MenuState:
		return m.menu.View()
	case TodoState:
		return m.todo.View()
	case TimerState:
		return m.timer.View()
	default:
		return "aruarian-tui\n\nunknown state\n"
	}
}
