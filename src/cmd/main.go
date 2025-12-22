package cmd

import (
	"log"
	"os"

	key "github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"kei/src/internal/ui/prompt"
)

type KeyMap struct {
	Quit   key.Binding
	Prompt key.Binding
}

var DefaultKeyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("quit", "quit kei"),
	),
	Prompt: key.NewBinding(
		key.WithKeys("/", ":"),
		key.WithHelp("prompt", "show prompt"),
	),
}

type RootModel struct {
	prompt     prompt.PromptModel
	showPrompt bool
}

func Run() {
	p := tea.NewProgram(RootModel{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.showPrompt {
		return m.prompt.Update(msg)
	}

	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch {
	case key.Matches(keyMsg, DefaultKeyMap.Quit):
		return m, tea.Quit

	case key.Matches(keyMsg, DefaultKeyMap.Prompt):
		if !m.showPrompt {
			m.showPrompt = true
			return m, nil
		} else {
			updatedPrompt, cmd := m.prompt.Update(msg)
			m.prompt = updatedPrompt.(prompt.PromptModel)
			return m, cmd
		}
	}

	return m, nil
}

func (m RootModel) View() string {
	if m.showPrompt {
		return m.prompt.View()
	}

	return ""
}
