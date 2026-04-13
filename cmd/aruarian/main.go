package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kajusviliusis/aruarian-tui/internal/app"
)

func main() {
	defer cleanup()
	prepareTerminal()

	p := tea.NewProgram(app.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func prepareTerminal() {
	fmt.Print("\033[2J\033[H")
	_ = os.Stdout.Sync()
}

func cleanup() {
	fmt.Print("\033[?25h")
	fmt.Print("\033[?1049l")
	fmt.Print("\033c")
	_ = os.Stdout.Sync()
}

