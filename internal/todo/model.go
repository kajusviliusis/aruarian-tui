package todo

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kajusviliusis/aruarian-tui/internal/styles"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

type Model struct {
	tasks      []Task
	cursor     int
	nextID     int
	storePath  string
	editing    bool
	editBuffer string
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

	if m.editing {
		return m.updateEditing(key)
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
			Title: "Task",
		})
		m.nextID++
		m.cursor = len(m.tasks) - 1
		mutated = true
	case "e":
		if len(m.tasks) > 0 {
			m.editing = true
			m.editBuffer = m.tasks[m.cursor].Title
		}
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

func (m Model) updateEditing(key tea.KeyMsg) (Model, tea.Cmd) {
	switch key.Type {
	case tea.KeyEsc:
		m.editing = false
		m.editBuffer = ""
	case tea.KeyEnter:
		if len(m.tasks) > 0 {
			title := strings.TrimSpace(m.editBuffer)
			if title != "" {
				m.tasks[m.cursor].Title = title
				_ = saveTasks(m.storePath, m.tasks)
			}
		}
		m.editing = false
		m.editBuffer = ""
	case tea.KeyBackspace:
		runes := []rune(m.editBuffer)
		if len(runes) > 0 {
			m.editBuffer = string(runes[:len(runes)-1])
		}
	case tea.KeySpace:
		m.editBuffer += " "
	case tea.KeyRunes:
		m.editBuffer += string(key.Runes)
	}

	return m, nil
}

func (m Model) Editing() bool {
	return m.editing
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString(styles.TitleStyle.Render("todo"))
	b.WriteString("\n\n")

	if len(m.tasks) == 0 {
		b.WriteString(styles.DimTextStyle.Render("no tasks yet. press 'a' to add one."))
		b.WriteString("\n")
	} else {
		for i, task := range m.tasks {
			prefix := "  "
			itemStyle := styles.ItemStyle

			if task.Completed {
				prefix = "  ✓ "
				itemStyle = styles.DimTextStyle
			}

			if i == m.cursor {
				if task.Completed {
					prefix = "> ✓ "
				} else {
					prefix = "> "
					itemStyle = styles.SelectedItemStyle
				}
			}

			b.WriteString(itemStyle.Render(prefix + task.Title))
			if i < len(m.tasks)-1 {
				b.WriteString("\n\n")
			} else {
				b.WriteString("\n")
			}
		}
	}

	b.WriteString("\n")
	if m.editing {
		b.WriteString(styles.DimTextStyle.Render("edit title:"))
		b.WriteString("\n")
		b.WriteString(styles.EditInputStyle.Render(m.editBuffer + "_"))
		b.WriteString("\n\n")
		b.WriteString(styles.DimTextStyle.Render("enter: save  esc: cancel  backspace: delete char"))
		return styles.ContainerStyle.Render(b.String())
	}

	b.WriteString(styles.DimTextStyle.Render("a: add  e: edit title  space: toggle  d: delete  esc: back"))
	return styles.ContainerStyle.Render(b.String())
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
