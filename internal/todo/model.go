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
	tasks     []Task
	cursor    int
	nextID    int
	storePath string
}

func NewModel(storePath string) Model {
	m := Model{
		nextID:    1,
		storePath: storePath,
	}

	tasks, err := loadTasks(storePath)
	if err == nil {
		m.tasks = tasks
		m.nextID = nextTaskID(tasks)
		if len(m.tasks) == 0 {
			m.cursor = 0
		} else if m.cursor >= len(m.tasks) {
			m.cursor = len(m.tasks) - 1
		}
	}

	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	mutated := false

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
		mutated = true
	case " ":
		if len(m.tasks) > 0 {
			m.tasks[m.cursor].Completed = !m.tasks[m.cursor].Completed
			mutated = true
		}
	case "d":
		if len(m.tasks) > 0 {
			m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor+1:]...)
			if m.cursor >= len(m.tasks) && m.cursor > 0 {
				m.cursor--
			}
			mutated = true
		}
	}

	if mutated {
		_ = saveTasks(m.storePath, m.tasks)
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

func nextTaskID(tasks []Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}
