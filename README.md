<p align="center">
  <h1 align="center">Gokedex</h1>
  <p align="center">
    <a href="https://github.com/knipferrc/gokedex/releases"><img src="https://img.shields.io/github/v/release/knipferrc/gokedex" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/knipferrc/gokedex?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/knipferrc/gokedex/actions"><img src="https://img.shields.io/github/workflow/status/knipferrc/gokedex/Release" alt="Build Status"></a>
  </p>
</p>

## About The Project

A pokedex for your terminal

### Built With

- [Go](https://golang.org/)
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- [bubbles](https://github.com/charmbracelet/bubbles)
- [lipgloss](https://github.com/charmbracelet/lipgloss)
- [Viper](https://github.com/spf13/viper)
- [Cobra](https://github.com/spf13/cobra)

## Installation

### Curl

```sh
curl -sfL https://raw.githubusercontent.com/knipferrc/gokedex/main/install.sh | sh
```

### Go

```
go install github.com/knipferrc/gokedex@latest
```

## Features

- Paginated list of pokemon
- Pokemon images shown as colored strings

## Configuration

- A config file will be generated at `~/.gokedex.yml` when you first run `gokedex`

```yml
settings:
  enable_logging: false
  enable_mousewheel: true
```

## Navigation

| Key        | Description                           |
| ---------- | ------------------------------------- |
| h or left  | Go to previous page of pokemon        |
| l or right | Go to next page of pokemon            |
| j or down  | Scroll the viewport down              |
| k or up    | Scroll the viewport up                |
| ?          | Toggle help screen                    |
| t          | Toggle between front and back sprites |
| q          | Quit                                  |
