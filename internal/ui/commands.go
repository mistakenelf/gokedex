package ui

import (
	"encoding/json"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/knipferrc/gokedex/internal/pokemon"

	tea "github.com/charmbracelet/bubbletea"
)

type pokemonMsg pokemon.Pokemon
type errMsg error

func (m Model) getPokemon(url string) tea.Cmd {
	return func() tea.Msg {
		apiUrl := ""
		if url == "" {
			apiUrl = "https://pokeapi.co/api/v2/pokemon?limit=5"
		} else {
			apiUrl = url
		}

		c := &http.Client{
			Timeout: 10 * time.Second,
		}

		// Get initial listing of pokemon.
		res, err := c.Get(apiUrl)
		if err != nil {
			return errMsg(err)
		}

		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var pokemonListResponse pokemon.Pokemon
		var pokemonDetailResponse pokemon.PokemonDetails
		var pokemonListing = make([]pokemon.PokemonDetails, 0)
		var stats = make([]pokemon.Stats, 0)

		json.Unmarshal(responseData, &pokemonListResponse)

		// Get details for each pokemon and add to pokemonDetailResponse.
		for _, pkmon := range pokemonListResponse.Results {
			pokemonDetails, err := c.Get(pkmon.URL)
			if err != nil {
				return errMsg(err)
			}

			pokemonDetailsData, err := ioutil.ReadAll(pokemonDetails.Body)
			if err != nil {
				return errMsg(err)
			}

			json.Unmarshal(pokemonDetailsData, &pokemonDetailResponse)

			//Get the front sprite of the pokemon.
			frontImageResponse, err := http.Get(pokemonDetailResponse.Sprites.FrontDefault)
			if err != nil {
				return errMsg(err)
			}

			// Decode the front sprite.
			m, err := png.Decode(frontImageResponse.Body)
			if err != nil {
				return errMsg(err)
			}

			// Convert the image to a string.
			pokemonFrontImage, _ := pokemon.ImageToString(20, 20, m)

			//Get the back sprite of the pokemon.
			backImageResponse, err := http.Get(pokemonDetailResponse.Sprites.BackDefault)
			if err != nil {
				return errMsg(err)
			}

			// Decode the front sprite.
			b, err := png.Decode(backImageResponse.Body)
			if err != nil {
				return errMsg(err)
			}

			// Convert the image to a string.
			pokemonBackImage, _ := pokemon.ImageToString(20, 20, b)

			pokemonListing = append(pokemonListing, pokemon.PokemonDetails{
				Name:    pokemonDetailResponse.Name,
				ID:      pokemonDetailResponse.ID,
				Sprites: pokemon.Sprites{FrontDefault: pokemonFrontImage, BackDefault: pokemonBackImage},
				Stats:   append(stats, pokemonDetailResponse.Stats...),
				Order:   pokemonDetailResponse.Order,
			})
		}

		pokemonList := pokemon.Pokemon{
			Count:    pokemonListResponse.Count,
			Next:     pokemonListResponse.Next,
			Previous: pokemonListResponse.Previous,
			Results:  pokemonListing,
		}

		return pokemonMsg(pokemonList)
	}
}
