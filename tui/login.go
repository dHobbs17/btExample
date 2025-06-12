package tui

import (
	"btexample/constants"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const LOGIN_PROMPT = "Please enter your username:"
const LOGIN_PLACEHOLDER = "username"

// MODEL DATA
type LoginMenu struct {
	textInput   textinput.Model
	err         error
	State       constants.State
	initialized bool
}

type errMsg error

func NewLoginModel() LoginMenu {
	ti := textinput.New()
	ti.Placeholder = LOGIN_PLACEHOLDER
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return LoginMenu{
		textInput: ti,
		err:       nil,
	}
}

func (m LoginMenu) Init() tea.Cmd {
	return textinput.Blink
}

// VIEW
func (s LoginMenu) View() string {
	return fmt.Sprintf(
		LOGIN_PROMPT+"\n\n%s\n\n%s",
		s.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

// UPDATE
func (s LoginMenu) Update(msg tea.Msg) (LoginMenu, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return s, tea.Quit

		case tea.KeyEnter:
			s.State = constants.ShowGame
			return s, nil
		}

	// We handle errors just like any other message
	case errMsg:
		s.err = msg
		return s, nil
	}

	s.textInput, cmd = s.textInput.Update(msg)
	return s, cmd
}
