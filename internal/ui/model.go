package ui

import (
	"github.com/knipferrc/gokedex/internal/constants"
	"github.com/knipferrc/gokedex/internal/pokemon"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	keys        keyMap
	help        help.Model
	loader      spinner.Model
	pokemon     pokemon.Model
	viewport    viewport.Model
	loadingMore bool
	err         error
	ready       bool
}

func NewModel() Model {
	keys := getDefaultKeyMap()

	l := spinner.NewModel()
	l.Spinner = spinner.Dot

	h := help.NewModel()
	h.Styles.FullKey.Foreground(lipgloss.Color(constants.White))
	h.Styles.FullDesc.Foreground(lipgloss.Color(constants.White))

	return Model{
		keys:        keys,
		help:        h,
		loader:      l,
		pokemon:     pokemon.Model{},
		viewport:    viewport.Model{},
		loadingMore: false,
		err:         nil,
		ready:       false,
	}
}
