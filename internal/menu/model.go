package menu

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kajusviliusis/aruarian-tui/internal/styles"
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

	b.WriteString(styles.Header.Render("aruarian-tui"))
	b.WriteString("\n\n")

	for i, item := range m.items {
		if i == m.cursor {
			b.WriteString(styles.MenuItemActive.Render("❯ " + item))
		} else {
			b.WriteString(styles.MenuItem.Render("  " + item))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(styles.Footer.Render("↑↓ or k/j: move  enter: select  q: quit"))

	return styles.Container.Render(b.String())
}

