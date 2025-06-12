package tui

import (
	"btexample/constants"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

const MAIN_MENU_GREETING = "Welcome to rpgcmd"
const MAIN_MENU_PROMPT = "-Press Enter to Login"

// MODEL DATA
type MainMenu struct {
	header      string
	prompt      string
	State       constants.State
	initialized bool
}

// REMOVE LATER
func NewMainModel() MainMenu {
	return MainMenu{header: MAIN_MENU_GREETING, prompt: MAIN_MENU_PROMPT}
}

func initMainMenu() MainMenu {
	return MainMenu{header: MAIN_MENU_GREETING, prompt: MAIN_MENU_PROMPT}
}

func (s MainMenu) Init() tea.Cmd {
	return nil
}

// VIEW
func (s MainMenu) View() string {
	textLen := len(s.header)
	topAndBottomBar := strings.Repeat("*", textLen+4)
	return fmt.Sprintf(
		"%s\n* %s *\n%s\n\n%s \n\nPress Ctrl+C or q to exit",
		topAndBottomBar, s.header, topAndBottomBar, s.prompt,
	)
}

// UPDATE
func (s MainMenu) Update(msg tea.Msg) (MainMenu, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			//fmt.Println("PRESSED ENTER")
			s.State = constants.ShowLogin
			return s, nil
		case "ctrl+c", "q":
			return s, tea.Quit
		}
	}

	return s, nil
}
