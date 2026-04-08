package app

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	state AppState
}

func NewModel() Model {
	return Model{state: MenuState}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	return "aruarian-tui\n\nStep 2: root state added (currently MenuState).\n\nPress q to quit.\n"
}
