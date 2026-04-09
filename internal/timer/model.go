package timer

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

type Model struct {
	defaultDuration time.Duration
	remaining       time.Duration
	running         bool
}

func NewModel(defaultDuration time.Duration) Model {
	if defaultDuration <= 0 {
		defaultDuration = 25 * time.Minute
	}

	return Model{
		defaultDuration: defaultDuration,
		remaining:       defaultDuration,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch typed := msg.(type) {
	case tea.KeyMsg:
		switch typed.String() {
		case "s":
			m.running = !m.running
			if m.running && m.remaining > 0 {
				return m, tickCmd()
			}
		case "r":
			m.running = false
			m.remaining = m.defaultDuration
		}
	case tickMsg:
		if !m.running {
			return m, nil
		}

		if m.remaining <= time.Second {
			m.remaining = 0
			m.running = false
			return m, nil
		}

		m.remaining -= time.Second
		return m, tickCmd()
	}

	return m, nil
}

func (m Model) Pause() Model {
	m.running = false
	return m
}

func (m Model) View() string {
	minutes := int(m.remaining / time.Minute)
	seconds := int((m.remaining % time.Minute) / time.Second)

	status := "paused"
	if m.running {
		status = "running"
	}

	return fmt.Sprintf(
		"TIMER\n\n%02d:%02d (%s)\n\ns: start/pause  r: reset  esc: back to menu\n",
		minutes,
		seconds,
		status,
	)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
