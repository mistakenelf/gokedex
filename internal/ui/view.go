package ui

import (
	"fmt"

	"github.com/knipferrc/gokedex/internal/constants"

	"github.com/charmbracelet/lipgloss"
)

// View returns a string representation of the entire application UI.
func (m Model) View() string {
	// If the viewport is not ready or we have no pokemon to display, return a spinner.
	if !m.ready || len(m.pokemon.Content.Results) == 0 || m.loadingMore {
		return fmt.Sprintf("%s%s", m.loader.View(), "loading...")
	}

	currentView := ""
	if m.help.ShowAll {
		currentView = m.help.View(m.keys)
	} else {
		currentView = m.viewport.View()
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(constants.White)).
		Bold(true).
		Italic(true).
		Render(currentView)
}
