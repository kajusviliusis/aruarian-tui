package todo

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

type Model struct {
	tasks  []Task
	cursor int
	nextID int
}

func NewModel() Model {
	return Model{nextID: 1}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.tasks)-1 {
			m.cursor++
		}
	case "a":
		m.tasks = append(m.tasks, Task{
			ID:    m.nextID,
			Title: fmt.Sprintf("Task %d", m.nextID),
		})
		m.nextID++
		m.cursor = len(m.tasks) - 1
	case " ":
		if len(m.tasks) > 0 {
			m.tasks[m.cursor].Completed = !m.tasks[m.cursor].Completed
		}
	case "d":
		if len(m.tasks) > 0 {
			m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor+1:]...)
			if m.cursor >= len(m.tasks) && m.cursor > 0 {
				m.cursor--
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString("TODO\n\n")

	if len(m.tasks) == 0 {
		b.WriteString("No tasks yet. Press 'a' to add one.\n")
	} else {
		for i, task := range m.tasks {
			prefix := "  "
			if i == m.cursor {
				prefix = "> "
			}

			status := "[ ]"
			if task.Completed {
				status = "[x]"
			}

			b.WriteString(fmt.Sprintf("%s%s %s\n", prefix, status, task.Title))
		}
	}

	b.WriteString("\n")
	b.WriteString("a: add  space: toggle  d: delete  up/down or k/j: move  esc: back to menu\n")
	return b.String()
}

