package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/kajusviliusis/aruarian-tui/internal/app"
)

func main() {
	p := tea.NewProgram(app.NewModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

