package ui

import "github.com/charmbracelet/bubbles/spinner"

type Pokemon struct {
	Count int `json:"count"`
}

type Model struct {
	loader  spinner.Model
	pokemon Pokemon
	err     error
	ready   bool
}

func NewModel() Model {
	return Model{
		loader:  spinner.NewModel(),
		pokemon: Pokemon{},
		err:     nil,
		ready:   false,
	}
}
