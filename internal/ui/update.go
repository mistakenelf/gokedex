package ui

import (
	"github.com/knipferrc/gokedex/internal/pokemon"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case pokemonMsg:
		m.pokemon.SetContent(pokemon.Pokemon(msg))
		m.viewport.SetContent(m.pokemon.View())
	case errMsg:
		m.err = msg
	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height
		m.viewport.Width = msg.Width
		m.viewport.SetContent(m.pokemon.View())

		if !m.ready {
			m.ready = true
		}

	case tea.KeyMsg:
		switch msg.String() {
		// Exit Gokedex.
		case "ctrl+c":
			return m, tea.Quit

		case "right":
			return m, m.loadNewPokemon(m.pokemon.Content.Next)

		case "left":
			return m, m.loadNewPokemon(m.pokemon.Content.Previous)
		}
	}

	m.loader, cmd = m.loader.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
