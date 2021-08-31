package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/knipferrc/gokedex/internal/pokemon"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pokemonMsg:
		m.pokemon.SetContent(pokemon.Pokemon(msg))
	case errMsg:
		m.err = msg
	case tea.WindowSizeMsg:
		if !m.ready {
			m.ready = true
		}

	case tea.KeyMsg:
		switch msg.String() {
		// Exit FM.
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}
