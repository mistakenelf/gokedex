package ui

import (
	"github.com/knipferrc/gokedex/internal/pokemon"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// keyMap struct contains all keybindings.
type keyMap struct {
	Help  key.Binding
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Quit  key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Up, k.Down, k.Left, k.Right, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Up, k.Down, k.Left, k.Right, k.Quit},
	}
}

type Model struct {
	keys     keyMap
	help     help.Model
	loader   spinner.Model
	pokemon  pokemon.Model
	viewport viewport.Model
	err      error
	ready    bool
}

func NewModel() Model {
	l := spinner.NewModel()
	l.Spinner = spinner.Dot

	h := help.NewModel()
	h.Styles.FullKey.Foreground(lipgloss.Color("#ffffff"))
	h.Styles.FullDesc.Foreground(lipgloss.Color("#ffffff"))

	// keys represents the key bindings in the app along with the help text.
	var keys = keyMap{
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "move left"),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "move right"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}

	return Model{
		keys:     keys,
		help:     h,
		loader:   l,
		pokemon:  pokemon.Model{},
		viewport: viewport.Model{},
		err:      nil,
		ready:    false,
	}
}
