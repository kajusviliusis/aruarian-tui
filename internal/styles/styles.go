package styles

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	BgColor      = lipgloss.Color("235")
	TitleColor   = lipgloss.Color("230")
	Accent       = lipgloss.Color("110")
	TextColor    = lipgloss.Color("252")
	DimColor     = lipgloss.Color("243")
	SuccessColor = lipgloss.Color("151")
)

var ContainerStyle = lipgloss.NewStyle().
	Padding(2, 4, 2, 4).
	Background(BgColor)

var TitleStyle = lipgloss.NewStyle().
	Foreground(TitleColor).
	Bold(false).
	Background(BgColor)

var SelectedItemStyle = lipgloss.NewStyle().
	Foreground(Accent).
	Bold(false).
	Background(BgColor)

var ItemStyle = lipgloss.NewStyle().
	Foreground(TextColor).
	Background(BgColor)

var DimTextStyle = lipgloss.NewStyle().
	Foreground(DimColor).
	Background(BgColor)

var TimerDisplayStyle = lipgloss.NewStyle().
	Foreground(Accent).
	Background(BgColor)

var ProgressFilledStyle = lipgloss.NewStyle().
	Foreground(Accent).
	Background(BgColor)

var ProgressEmptyStyle = lipgloss.NewStyle().
	Foreground(DimColor).
	Background(BgColor)

var SuccessStyle = lipgloss.NewStyle().
	Foreground(SuccessColor).
	Background(BgColor)

var EditInputStyle = lipgloss.NewStyle().
	Foreground(Accent).
	Background(BgColor)

var EditCursorStyle = lipgloss.NewStyle().
	Foreground(Accent).
	Background(BgColor)

var (
	Container           = ContainerStyle
	Header              = TitleStyle
	MenuItem            = ItemStyle
	MenuItemActive      = SelectedItemStyle
	TaskItem            = ItemStyle
	TaskItemActive      = SelectedItemStyle
	TaskCompleted       = DimTextStyle
	TaskCompletedActive = DimTextStyle
	CheckboxUnchecked   = DimTextStyle
	CheckboxChecked     = SuccessStyle
	EditInput           = EditInputStyle
	EditCursor          = EditCursorStyle
	Footer              = DimTextStyle
	StatusRunning       = SuccessStyle
	StatusPaused        = DimTextStyle
	TimerDisplay        = TimerDisplayStyle
)

var bigDigitGlyphs = map[rune][]string{
	'0': {"████ ", "█  █ ", "█  █ ", "█  █ ", "████ "},
	'1': {"  ██ ", " ███ ", "  ██ ", "  ██ ", "█████"},
	'2': {"████ ", "   █ ", " ███ ", "█    ", "█████"},
	'3': {"████ ", "   █ ", " ███ ", "   █ ", "████ "},
	'4': {"█  █ ", "█  █ ", "█████", "   █ ", "   █ "},
	'5': {"█████", "█    ", "████ ", "   █ ", "████ "},
	'6': {"████ ", "█    ", "████ ", "█  █ ", "████ "},
	'7': {"█████", "   █ ", "  █  ", " █   ", "█    "},
	'8': {"████ ", "█  █ ", "████ ", "█  █ ", "████ "},
	'9': {"████ ", "█  █ ", "█████", "   █ ", "████ "},
	':': {"     ", "  █  ", "     ", "  █  ", "     "},
	' ': {"     ", "     ", "     ", "     ", "     "},
}

func CenterContent(content string, width, height int) string {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Background(BgColor).
		Render(content)
}

func CenterLines(lines ...string) string {
	maxWidth := 0
	for _, line := range lines {
		if width := lipgloss.Width(line); width > maxWidth {
			maxWidth = width
		}
	}

	if maxWidth == 0 {
		return strings.Join(lines, "\n")
	}

	centered := make([]string, len(lines))
	lineStyle := lipgloss.NewStyle().
		Width(maxWidth).
		Align(lipgloss.Center).
		Background(BgColor)
	for i, line := range lines {
		centered[i] = lineStyle.Render(line)
	}

	return strings.Join(centered, "\n")
}

func RenderBigDigits(text string) string {
	const height = 5
	if text == "" {
		return ""
	}

	rows := make([]strings.Builder, height)
	runes := []rune(text)
	for idx, r := range runes {
		glyph, ok := bigDigitGlyphs[r]
		if !ok {
			glyph = bigDigitGlyphs[' ']
		}
		for i := 0; i < height; i++ {
			rows[i].WriteString(glyph[i])
			if idx < len(runes)-1 {
				rows[i].WriteString(" ")
			}
		}
	}

	lines := make([]string, height)
	for i := 0; i < height; i++ {
		lines[i] = TimerDisplayStyle.Render(rows[i].String())
	}

	return strings.Join(lines, "\n")
}

func ProgressBar(percentage float64, width int) string {
	if width <= 0 {
		width = 18
	}

	if percentage < 0 {
		percentage = 0
	}
	if percentage > 1 {
		percentage = 1
	}

	filled := int(float64(width)*percentage + 0.5)
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}

	var b strings.Builder
	b.WriteString(DimTextStyle.Render("◖"))
	for i := 0; i < width; i++ {
		if i < filled {
			b.WriteString(ProgressFilledStyle.Render("▰"))
		} else {
			b.WriteString(ProgressEmptyStyle.Render("▱"))
		}
	}
	b.WriteString(DimTextStyle.Render("◗"))

	return b.String()
}
