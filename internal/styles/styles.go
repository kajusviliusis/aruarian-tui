package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Accent = lipgloss.Color("6")

	Gray = lipgloss.Color("8")
	DarkGray = lipgloss.Color("0")
	LightGray = lipgloss.Color("7")

	Success = lipgloss.Color("2")
	Muted = lipgloss.Color("4")
	BgColor = lipgloss.Color("235")
)

var Container = lipgloss.NewStyle().
	Padding(1, 2).
	Margin(0).
	Background(BgColor)

var Header = lipgloss.NewStyle().
	Foreground(Accent).
	Bold(true).
	MarginBottom(1).
	Background(BgColor)

var MenuItem = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(LightGray).
	Background(BgColor)

var MenuItemActive = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(Accent).
	Bold(true).
	Background(DarkGray)

var TaskItem = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(LightGray).
	Background(BgColor)

var TaskItemActive = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(Accent).
	Background(DarkGray)

var TaskCompleted = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(Gray).
	Strikethrough(true).
	Background(BgColor)

var TaskCompletedActive = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(Gray).
	Background(DarkGray).
	Strikethrough(true)

var CheckboxUnchecked = lipgloss.NewStyle().
	Foreground(Gray).
	Background(BgColor)

var CheckboxChecked = lipgloss.NewStyle().
	Foreground(Success).
	Background(BgColor)

var EditInput = lipgloss.NewStyle().
	Foreground(Accent).
	Bold(true).
	Background(BgColor)

var EditCursor = lipgloss.NewStyle().
	Foreground(Accent).
	Blink(true).
	Background(BgColor)

var Footer = lipgloss.NewStyle().
	Foreground(Gray).
	MarginTop(1).
	Background(BgColor)

var StatusRunning = lipgloss.NewStyle().
	Foreground(Success).
	Bold(true).
	Background(BgColor)

var StatusPaused = lipgloss.NewStyle().
	Foreground(Gray).
	Background(BgColor)

var TimerDisplay = lipgloss.NewStyle().
	Foreground(Accent).
	Bold(true).
	Align(lipgloss.Center).
	MarginBottom(1).
	Background(BgColor)

func CenterContent(content string, width, height int) string {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Background(BgColor).
		Render(content)
}


