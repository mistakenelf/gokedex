package ui

import (
	"github.com/knipferrc/gokedex/internal/config"
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
	appConfig   config.Config
	loadingMore bool
	err         error
	ready       bool
}

func NewModel() Model {
	keys := getDefaultKeyMap()
	cfg := config.GetConfig()

	l := spinner.NewModel()
	l.Spinner = spinner.Dot

	h := help.NewModel()
	h.Styles.FullKey.Foreground(lipgloss.Color("#ffffff"))
	h.Styles.FullDesc.Foreground(lipgloss.Color("#ffffff"))

	return Model{
		keys:        keys,
		help:        h,
		loader:      l,
		pokemon:     pokemon.Model{},
		viewport:    viewport.Model{},
		appConfig:   cfg,
		loadingMore: false,
		err:         nil,
		ready:       false,
	}
}
