package tui

import (
	"btexample/constants"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	state constants.State
	menu  MainMenu
	login LoginMenu
	game  GameModel
}

func NewProject() model {
	return model{0, NewMainModel(), NewLoginModel(), NewGameModel()}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case constants.ShowMenu:
		var cmd tea.Cmd
		m.menu, cmd = m.menu.Update(msg)
		m.state = m.menu.State
		// Potential logic for changing the state goes here. You change the
		// state based on how the update affected the menu model.
		return m, cmd

	case constants.ShowLogin:
		var cmd tea.Cmd
		m.login, cmd = m.login.Update(msg)
		m.state = m.login.State
		// Potential logic for changing the state goes here. You change the
		// state based on how the update affected the section model.
		return m, cmd

	case constants.ShowGame:
		var cmd tea.Cmd
		m.game, cmd = m.game.Update(msg)
		m.state = m.game.State
		// Potential logic for changing the state goes here. You change the
		// state based on how the update affected the section model.
		return m, cmd
	default:
		return m, nil
	}
}

func (m model) View() string {
	switch m.state {
	case constants.ShowLogin:
		return m.login.View()
	case constants.ShowGame:
		return m.game.View()
	default:
		return m.menu.View()
	}
}
