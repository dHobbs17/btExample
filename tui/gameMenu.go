package tui

import (
	"btexample/constants"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type GameModel struct {
	choices     []string
	cursor      int
	selected    map[int]struct{}
	State       constants.State
	initialized bool
}

func NewGameModel() GameModel {
	return GameModel{
		choices:  []string{"Buy carrots", "Sell carrots"},
		selected: make(map[int]struct{}),
	}
}

func InitGameMenu() GameModel {
	return GameModel{
		choices:  []string{"Buy carrots", "Sell carrots"},
		selected: make(map[int]struct{}),
	}
}

func (m GameModel) Init() tea.Cmd {
	return nil
}

func (m GameModel) Update(msg tea.Msg) (GameModel, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m GameModel) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
