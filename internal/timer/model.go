package timer

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kajusviliusis/aruarian-tui/internal/styles"
)

const (
	minDuration  = 1 * time.Minute
	stepDuration = 1 * time.Minute
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
		case "+", "=":
			if m.running {
				return m, nil
			}
			m.defaultDuration += stepDuration
			m.remaining = m.defaultDuration
		case "-":
			if m.running {
				return m, nil
			}
			if m.defaultDuration > minDuration {
				m.defaultDuration -= stepDuration
				if m.defaultDuration < minDuration {
					m.defaultDuration = minDuration
				}
			}
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
	hours := int(m.remaining / time.Hour)
	minutes := int((m.remaining % time.Hour) / time.Minute)
	seconds := int((m.remaining % time.Minute) / time.Second)
	configuredMinutes := int(m.defaultDuration / time.Minute)

	timeDisplay := fmt.Sprintf("%02d:%02d", minutes, seconds)
	if hours > 0 {
		timeDisplay = fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}

	status := styles.StatusPaused.Render("paused")
	if m.running {
		status = styles.StatusRunning.Render("running")
	}

	var b strings.Builder
	b.WriteString(styles.Header.Render("TIMER"))
	b.WriteString("\n\n")
	b.WriteString(styles.TimerDisplay.Render(timeDisplay))
	b.WriteString("\n")
	b.WriteString(styles.MenuItem.Render(fmt.Sprintf("status: %s", status)))
	b.WriteString("\n")
	b.WriteString(styles.MenuItem.Render(fmt.Sprintf("configured: %dm", configuredMinutes)))
	b.WriteString("\n\n")
	b.WriteString(styles.Footer.Render("s: start/pause  r: reset  +/-: duration (paused)  esc: back to menu"))

	return styles.Container.Render(b.String())
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
