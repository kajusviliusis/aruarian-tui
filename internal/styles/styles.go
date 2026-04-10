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
)

var Container = lipgloss.NewStyle().
	Padding(1, 2).
	Margin(0)

var Header = lipgloss.NewStyle().
	Foreground(Accent).
	Bold(true).
	MarginBottom(1)

var MenuItem = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(LightGray)

var MenuItemActive = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(Accent).
	Bold(true).
	Background(DarkGray)

var TaskItem = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(LightGray)

var TaskItemActive = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(Accent).
	Background(DarkGray)

var TaskCompleted = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(Gray).
	Strikethrough(true)

var TaskCompletedActive = lipgloss.NewStyle().
	Padding(0, 1).
	Foreground(Gray).
	Background(DarkGray).
	Strikethrough(true)

var CheckboxUnchecked = lipgloss.NewStyle().
	Foreground(Gray)

var CheckboxChecked = lipgloss.NewStyle().
	Foreground(Success)

var EditInput = lipgloss.NewStyle().
	Foreground(Accent).
	Bold(true)

var EditCursor = lipgloss.NewStyle().
	Foreground(Accent).
	Blink(true)

var Footer = lipgloss.NewStyle().
	Foreground(Gray).
	MarginTop(1)

var StatusRunning = lipgloss.NewStyle().
	Foreground(Success).
	Bold(true)

var StatusPaused = lipgloss.NewStyle().
	Foreground(Gray)

var TimerDisplay = lipgloss.NewStyle().
	Foreground(Accent).
	Bold(true).
	Align(lipgloss.Center).
	MarginBottom(1)


