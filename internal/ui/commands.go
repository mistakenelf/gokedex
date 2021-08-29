package ui

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type pokemonMsg Pokemon
type errMsg struct{ error }

func loadPokemon() tea.Msg {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := c.Get("https://pokeapi.co/api/v2/pokemon?limit=20")
	if err != nil {
		return errMsg{err}
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Pokemon
	json.Unmarshal(responseData, &responseObject)

	return pokemonMsg(responseObject)
}
