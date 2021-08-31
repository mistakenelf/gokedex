package ui

import (
	"github.com/knipferrc/gokedex/internal/pokemon"

	"github.com/charmbracelet/bubbles/spinner"
)

type Model struct {
	loader  spinner.Model
	pokemon pokemon.Model
	err     error
	ready   bool
}

func NewModel() Model {
	return Model{
		loader:  spinner.NewModel(),
		pokemon: pokemon.Model{},
		err:     nil,
		ready:   false,
	}
}
