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
	tasks       []Task
	cursor      int
	nextID      int
	storePath   string
	editing     bool
	editingNew  bool
	editBuffer  string
	editCursor  int
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
		m.tasks = append(m.tasks, Task{ID: m.nextID})
		m.nextID++
		m.cursor = len(m.tasks) - 1
		m.editing = true
		m.editingNew = true
		m.editBuffer = ""
		m.editCursor = 0
	case "e":
		if len(m.tasks) > 0 {
			m.editing = true
			m.editingNew = false
			m.editBuffer = m.tasks[m.cursor].Title
			m.editCursor = len([]rune(m.editBuffer))
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
		if m.editingNew {
			m.removeDraftTask()
			_ = saveTasks(m.storePath, m.tasks)
		}
		m.editing = false
		m.editingNew = false
		m.editBuffer = ""
		m.editCursor = 0
	case tea.KeyEnter:
		title := strings.TrimSpace(m.editBuffer)
		if m.editingNew {
			if title == "" {
				m.removeDraftTask()
			} else if len(m.tasks) > 0 {
				m.tasks[m.cursor].Title = title
			}
			_ = saveTasks(m.storePath, m.tasks)
		} else if title != "" && len(m.tasks) > 0 {
			m.tasks[m.cursor].Title = title
			_ = saveTasks(m.storePath, m.tasks)
		}
		m.editing = false
		m.editingNew = false
		m.editBuffer = ""
		m.editCursor = 0
	case tea.KeyLeft:
		if m.editCursor > 0 {
			m.editCursor--
		}
	case tea.KeyRight:
		if m.editCursor < len([]rune(m.editBuffer)) {
			m.editCursor++
		}
	case tea.KeyBackspace:
		if m.editCursor > 0 {
			runes := []rune(m.editBuffer)
			idx := m.editCursor - 1
			runes = append(runes[:idx], runes[idx+1:]...)
			m.editBuffer = string(runes)
			m.editCursor--
		}
	case tea.KeyDelete:
		runes := []rune(m.editBuffer)
		if m.editCursor < len(runes) {
			runes = append(runes[:m.editCursor], runes[m.editCursor+1:]...)
			m.editBuffer = string(runes)
		}
	case tea.KeySpace:
		m.insertEditRune(' ')
	case tea.KeyRunes:
		for _, r := range key.Runes {
			m.insertEditRune(r)
		}
	}

	return m, nil
}

func (m *Model) insertEditRune(r rune) {
	runes := []rune(m.editBuffer)
	idx := m.editCursor
	runes = append(runes[:idx], append([]rune{r}, runes[idx:]...)...)
	m.editBuffer = string(runes)
	m.editCursor++
}

func (m *Model) removeDraftTask() {
	if !m.editingNew || len(m.tasks) == 0 {
		return
	}

	m.tasks = append(m.tasks[:m.cursor], m.tasks[m.cursor+1:]...)
	if len(m.tasks) == 0 {
		m.cursor = 0
		return
	}
	if m.cursor >= len(m.tasks) {
		m.cursor = len(m.tasks) - 1
	}
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
		b.WriteString(styles.EditInputStyle.Render(renderEditInput(m.editBuffer, m.editCursor)))
		b.WriteString("\n\n")
		b.WriteString(styles.DimTextStyle.Render("enter: save  esc: cancel  backspace: delete char"))
		return styles.ContainerStyle.Render(b.String())
	}

	b.WriteString(styles.DimTextStyle.Render("a: add  e: edit title  space: toggle  d: delete  esc: back"))
	return styles.ContainerStyle.Render(b.String())
}

func renderEditInput(buffer string, cursor int) string {
	runes := []rune(buffer)
	if cursor < 0 {
		cursor = 0
	}
	if cursor > len(runes) {
		cursor = len(runes)
	}

	left := string(runes[:cursor])
	right := string(runes[cursor:])
	return left + "_" + right
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
