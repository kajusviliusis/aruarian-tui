package menu

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kajusviliusis/aruarian-tui/internal/styles"
)

type Selection string

const (
	SelectionNotes Selection = "notes"
	SelectionTodo  Selection = "todo"
	SelectionTimer Selection = "timer"
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
			case "notes":
				return m, func() tea.Msg { return SelectionMsg{Selection: SelectionNotes} }
			case "todo":
				return m, func() tea.Msg { return SelectionMsg{Selection: SelectionTodo} }
			case "timer":
				return m, func() tea.Msg { return SelectionMsg{Selection: SelectionTimer} }
			case "quit":
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("\n\n")
	b.WriteString(styles.TitleStyle.Render(
		"                             _             \n" +
			"  __ _ _ __ _   _  __ _ _ __(_) __ _ _ __  \n" +
			" / _` | '__| | | |/ _` | '__| |/ _` | '_ \\ \n" +
			"| (_| | |  | |_| | (_| | |  | | (_| | | | |\n" +
			" \\__,_|_|   \\__,_|\\__,_|_|  |_|\\__,_|_| |_|",
	))
	b.WriteString("\n\n")

	for i, item := range m.items {
		line := "  " + item
		itemStyle := styles.ItemStyle
		if i == m.cursor {
			line = "> " + item
			itemStyle = styles.SelectedItemStyle
		}

		b.WriteString(itemStyle.Render(line))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(styles.DimTextStyle.Render("calm focus."))
	b.WriteString("\n\n")
	b.WriteString(styles.DimTextStyle.Render("↑↓ or k/j  enter  q"))

	return styles.ContainerStyle.Render(b.String())
}
