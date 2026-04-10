package app

import (
	"os"
	"path/filepath"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kajusviliusis/aruarian-tui/internal/menu"
	"github.com/kajusviliusis/aruarian-tui/internal/notes"
	"github.com/kajusviliusis/aruarian-tui/internal/styles"
	"github.com/kajusviliusis/aruarian-tui/internal/timer"
	"github.com/kajusviliusis/aruarian-tui/internal/todo"
)

const notesLaunchDelay = 120 * time.Millisecond

type notesLaunchMsg struct{}

type Model struct {
	state  AppState
	menu   menu.Model
	timer  timer.Model
	todo   todo.Model
	width  int
	height int
}

func NewModel() Model {
	return Model{
		state: MenuState,
		menu: menu.NewModel([]string{
			"notes",
			"todo",
			"timer",
			"quit",
		}),
		timer: timer.NewModel(1 * time.Hour),
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
	case tea.WindowSizeMsg:
		m.width = typed.Width
		m.height = typed.Height
	case menu.SelectionMsg:
		switch typed.Selection {
		case menu.SelectionNotes:
			m.state = NotesState
			return m, tea.Tick(notesLaunchDelay, func(time.Time) tea.Msg { return notesLaunchMsg{} })
		case menu.SelectionTodo:
			m.state = TodoState
			return m, nil
		case menu.SelectionTimer:
			m.state = TimerState
			return m, nil
		}
	case notesLaunchMsg:
		return m, notes.LaunchNeovim()
	case notes.ExitMsg:
		m.state = MenuState
		return m, nil
	}

	switch m.state {
	case MenuState:
		var cmd tea.Cmd
		m.menu, cmd = m.menu.Update(msg)
		return m, cmd
	case TodoState:
		if key, ok := msg.(tea.KeyMsg); ok && key.String() == "esc" && !m.todo.Editing() {
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
	case NotesState:
		return m, nil
	default:
		return m, nil
	}
}

func (m Model) View() string {
	switch m.state {
	case MenuState:
		return styles.CenterContent(m.menu.View(), m.width, m.height)
	case TodoState:
		return styles.CenterContent(m.todo.View(), m.width, m.height)
	case TimerState:
		return styles.CenterContent(m.timer.View(), m.width, m.height)
	case NotesState:
		return styles.CenterContent(styles.ContainerStyle.Render(styles.TitleStyle.Render("opening notes...")), m.width, m.height)
	default:
		return styles.CenterContent("unknown state\n", m.width, m.height)
	}
}
