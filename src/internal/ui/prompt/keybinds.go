package prompt

import (
	key "github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyMap struct {
	Insert    key.Binding
	Normal    key.Binding
	Backspace key.Binding
	Quit      key.Binding
}

var DefaultKeyMap = KeyMap{
	Insert: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("insert", "insert"),
	),
	Normal: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("normal", "normal"),
	),
	Backspace: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "backspace"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c", "esc"),
		key.WithHelp("quit", "quit kei"),
	),
}

func (m PromptModel) HandleKeybindings(msg tea.Msg) (tea.Model, tea.Cmd) {
	keyMsg, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	if !m.focus {
		switch {
		case key.Matches(keyMsg, DefaultKeyMap.Insert):
			m.focus = true
			return m, nil

		case key.Matches(keyMsg, DefaultKeyMap.Quit):
			return m, tea.Quit
		}

		return m, nil
	}

	switch {
	case key.Matches(keyMsg, DefaultKeyMap.Normal):
		m.focus = false
		return m, nil

	case key.Matches(keyMsg, DefaultKeyMap.Backspace):
		if len(m.input) > 0 {
			m.input = m.input[:len(m.input)-1]
		}
		return m, nil

	case keyMsg.Type == tea.KeyRunes:
		m.input += string(keyMsg.Runes)
	}

	return m, nil
}
