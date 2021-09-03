package ui

import (
	"encoding/json"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/knipferrc/gokedex/internal/constants"
	"github.com/knipferrc/gokedex/internal/pokemon"

	tea "github.com/charmbracelet/bubbletea"
)

type pokemonMsg pokemon.Pokemon
type errMsg struct{ error }

func loadInitialPokemonData() tea.Msg {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := c.Get(constants.ApiUrl)
	if err != nil {
		return errMsg{err}
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject pokemon.Pokemon
	var detailResponseObject pokemon.PokemonDetails
	var finalDetails = make([]pokemon.PokemonDetails, 0)
	var stats = make([]pokemon.Stats, 0)

	json.Unmarshal(responseData, &responseObject)

	// Get details for each pokemon and add to responseObject.
	for _, pkmon := range responseObject.Results {
		pokemonDetails, err := c.Get(pkmon.URL)
		if err != nil {
			return errMsg{err}
		}

		pokemonDetailsData, err := ioutil.ReadAll(pokemonDetails.Body)
		if err != nil {
			return errMsg{err}
		}

		json.Unmarshal(pokemonDetailsData, &detailResponseObject)

		response, err := http.Get(detailResponseObject.Sprites.FrontDefault)
		if err != nil {
			return errMsg{err}
		}

		m, err := png.Decode(response.Body)
		if err != nil {
			return errMsg{err}
		}

		pokemonImageString, _ := pokemon.ImageToString(20, 20, m)

		finalDetails = append(finalDetails, pokemon.PokemonDetails{
			Name:    detailResponseObject.Name,
			ID:      detailResponseObject.ID,
			Sprites: pokemon.Sprites{FrontDefault: pokemonImageString},
			Stats:   append(stats, detailResponseObject.Stats...),
			Order:   detailResponseObject.Order,
		})
	}

	finalRes := pokemon.Pokemon{
		Count:    responseObject.Count,
		Next:     responseObject.Next,
		Previous: responseObject.Previous,
		Results:  finalDetails,
	}

	return pokemonMsg(finalRes)
}

func (m Model) loadNewPokemon(url string) tea.Cmd {
	return func() tea.Msg {
		apiUrl := ""
		if url == "" {
			apiUrl = constants.ApiUrl
		} else {
			apiUrl = url
		}

		c := &http.Client{
			Timeout: 10 * time.Second,
		}

		res, err := c.Get(apiUrl)
		if err != nil {
			return errMsg{err}
		}

		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var responseObject pokemon.Pokemon
		var detailResponseObject pokemon.PokemonDetails
		var finalDetails = make([]pokemon.PokemonDetails, 0)
		var stats = make([]pokemon.Stats, 0)

		json.Unmarshal(responseData, &responseObject)

		// Get details for each pokemon and add to responseObject.
		for _, pkmon := range responseObject.Results {
			pokemonDetails, err := c.Get(pkmon.URL)
			if err != nil {
				return errMsg{err}
			}

			pokemonDetailsData, err := ioutil.ReadAll(pokemonDetails.Body)
			if err != nil {
				return errMsg{err}
			}

			json.Unmarshal(pokemonDetailsData, &detailResponseObject)

			response, err := http.Get(detailResponseObject.Sprites.FrontDefault)
			if err != nil {
				return errMsg{err}
			}

			m, err := png.Decode(response.Body)
			if err != nil {
				return errMsg{err}
			}

			pokemonImageString, _ := pokemon.ImageToString(20, 20, m)

			finalDetails = append(finalDetails, pokemon.PokemonDetails{
				Name:    detailResponseObject.Name,
				ID:      detailResponseObject.ID,
				Sprites: pokemon.Sprites{FrontDefault: pokemonImageString},
				Stats:   append(stats, detailResponseObject.Stats...),
				Order:   detailResponseObject.Order,
			})
		}

		finalRes := pokemon.Pokemon{
			Count:    responseObject.Count,
			Next:     responseObject.Next,
			Previous: responseObject.Previous,
			Results:  finalDetails,
		}

		return pokemonMsg(finalRes)
	}
}
