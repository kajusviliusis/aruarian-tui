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
		defaultDuration = 1 * time.Hour
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

	status := "paused"
	statusStyle := styles.DimTextStyle
	if m.running {
		status = "running"
		statusStyle = styles.SuccessStyle
	}

	progress := 0.0
	if m.defaultDuration > 0 {
		progress = float64(m.defaultDuration-m.remaining) / float64(m.defaultDuration)
	}

	digits := strings.Split(styles.RenderBigDigits(timeDisplay), "\n")
	progressLine := styles.ProgressBar(progress, 24)
	percentageLine := styles.DimTextStyle.Render(fmt.Sprintf("%02.0f%%", progress*100))

	lines := []string{
		//styles.TitleStyle.Render("deep work"),
		"",
	}
	lines = append(lines, digits...)
	lines = append(lines,
		"",
		styles.DimTextStyle.Render("remaining"),
		progressLine,
		percentageLine,
		"",
		statusStyle.Render(fmt.Sprintf("%s  •  %dm", status, configuredMinutes)),
		"",
		styles.DimTextStyle.Render("s: start/pause   r: reset   esc: back   +/-: adjust (paused)"),
	)

	return styles.ContainerStyle.Render(styles.CenterLines(lines...))
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
