package ui

import (
	"github.com/knipferrc/gokedex/internal/pokemon"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
)

type Model struct {
	loader   spinner.Model
	pokemon  pokemon.Model
	viewport viewport.Model
	err      error
	ready    bool
}

func NewModel() Model {
	l := spinner.NewModel()
	l.Spinner = spinner.Dot

	return Model{
		loader:   l,
		pokemon:  pokemon.Model{},
		viewport: viewport.Model{},
		err:      nil,
		ready:    false,
	}
}
