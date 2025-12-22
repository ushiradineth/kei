package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type PromptModel struct {
	input string
	focus bool
}

func (m PromptModel) Init() tea.Cmd {
	return nil
}

func (m PromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.HandleKeybindings(msg)
	}

	return m, nil
}

func (m PromptModel) View() string {
	if m.focus {
		return fmt.Sprintf("Input: %s", m.input)
	} else {
		return fmt.Sprintf("Input: %s █", m.input)
	}
}
