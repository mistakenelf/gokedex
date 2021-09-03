package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/knipferrc/gokedex/internal/constants"
	"github.com/knipferrc/gokedex/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gokedex",
	Short:   "Gokedex is a pokedex for your terminal",
	Version: constants.AppVersion,
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m := ui.NewModel()
		var opts []tea.ProgramOption

		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		defer func() {
			if err = f.Close(); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		}()

		// Always append alt screen program option.
		opts = append(opts, tea.WithAltScreen())

		// Append mouse support.
		opts = append(opts, tea.WithMouseAllMotion())

		// Initialize new app.
		p := tea.NewProgram(m, opts...)
		if err := p.Start(); err != nil {
			log.Fatal("Failed to start gokedex", err)
			os.Exit(1)
		}
	},
}

// Execute executes the root command which starts the application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
