package ui

import (
	"github.com/knipferrc/gokedex/internal/pokemon"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case pokemonMsg:
		m.loadingMore = false
		m.pokemon.SetContent(pokemon.Pokemon(msg))
		m.viewport.SetContent(m.pokemon.View())
	case errMsg:
		m.err = msg
	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height
		m.viewport.Width = msg.Width
		m.help.Width = msg.Width
		m.viewport.SetContent(m.pokemon.View())

		if !m.ready {
			m.pokemon.ToggleImage(false)
			m.ready = true
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

			return m, nil
		case key.Matches(msg, m.keys.Right):
			m.loadingMore = true
			return m, m.getPokemon(m.pokemon.Content.Next)
		case key.Matches(msg, m.keys.Left):
			m.loadingMore = true
			return m, m.getPokemon(m.pokemon.Content.Previous)
		case key.Matches(msg, m.keys.ToggleImage):
			m.pokemon.ToggleImage(!m.pokemon.ShowBack)
			m.viewport.SetContent(m.pokemon.View())
			return m, nil
		}
	}

	m.loader, cmd = m.loader.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
