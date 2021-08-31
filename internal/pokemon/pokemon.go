package pokemon

import (
	"fmt"
	"image"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/nfnt/resize"
)

type Stat struct {
	Name string `json:"name"`
}

type Stats struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     Stat
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
}

type PokemonDetails struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Order   int    `json:"order"`
	Sprites Sprites
	Stats   []Stats
}

type Pokemon struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []PokemonDetails
}

type Model struct {
	Content Pokemon
}

func (m *Model) SetContent(content Pokemon) {
	m.Content = content
}

// ImageToString converts an image to a string.
func ImageToString(width, height uint, img image.Image) (string, error) {
	img = resize.Thumbnail(width, height*2-4, img, resize.Lanczos3)
	b := img.Bounds()
	w := b.Max.X
	h := b.Max.Y
	str := strings.Builder{}
	for y := 0; y < h; y += 2 {
		for x := w; x < int(width); x = x + 2 {
			str.WriteString(" ")
		}
		for x := 0; x < w; x++ {
			c1, _ := colorful.MakeColor(img.At(x, y))
			color1 := lipgloss.Color(c1.Hex())
			c2, _ := colorful.MakeColor(img.At(x, y+1))
			color2 := lipgloss.Color(c2.Hex())
			str.WriteString(lipgloss.NewStyle().Foreground(color1).Background(color2).Render("▀"))
		}
		str.WriteString("\n")
	}

	return str.String(), nil
}

func (m Model) View() string {
	pokemonList := ""

	for _, pokemon := range m.Content.Results {
		pokemonList += fmt.Sprintf("%s\n%s\n", pokemon.Name, pokemon.Sprites.FrontDefault)
	}

	return pokemonList
}
