package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, spinner.Tick)
	cmds = append(cmds, loadInitialPokemonData)

	return tea.Batch(cmds...)
}
