package menu

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Selection string

const (
	SelectionNotes Selection = "NOTES"
	SelectionTodo  Selection = "TODO"
	SelectionTimer Selection = "TIMER"
)

type SelectionMsg struct {
	Selection Selection
}

type Model struct {
	items  []string
	cursor int
}

func NewModel(items []string) Model {
	return Model{items: items}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.items) == 0 {
				return m, nil
			}

			switch m.items[m.cursor] {
			case "NOTES":
				return m, func() tea.Msg { return SelectionMsg{Selection: SelectionNotes} }
			case "TODO":
				return m, func() tea.Msg { return SelectionMsg{Selection: SelectionTodo} }
			case "TIMER":
				return m, func() tea.Msg { return SelectionMsg{Selection: SelectionTimer} }
			case "QUIT":
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString("aruarian-tui\n\n")

	for i, item := range m.items {
		prefix := "  "
		if i == m.cursor {
			prefix = "> "
		}
		b.WriteString(fmt.Sprintf("%s%s\n", prefix, item))
	}

	b.WriteString("\n")
	b.WriteString("up/down or k/j: move  enter: select  q: quit\n")
	return b.String()
}

