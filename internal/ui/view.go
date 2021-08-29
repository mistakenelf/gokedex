package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// View returns a string representation of the entire application UI.
func (m Model) View() string {
	// If the viewport is not ready, return a spinner.
	if !m.ready {
		return fmt.Sprintf("%s%s", m.loader.View(), "loading...")
	}

	return lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Bold(true).Render(fmt.Sprintf("%d", m.pokemon.Count))
}
