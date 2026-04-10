package notes

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type ExitMsg struct{}

func LaunchNeovim() tea.Cmd {
	return tea.ExecProcess(createNvimCmd(), func(err error) tea.Msg {
		return ExitMsg{}
	})
}

func createNvimCmd() *exec.Cmd {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	notesDir := homeDir + "/notes"
	_ = os.MkdirAll(notesDir, 0o755)

	cmd := exec.Command("nvim")
	cmd.Dir = notesDir
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
