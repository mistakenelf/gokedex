package ui

import "github.com/charmbracelet/bubbles/spinner"

type Model struct {
	loader spinner.Model
	ready  bool
}

func NewModel() Model {
	return Model{}
}
