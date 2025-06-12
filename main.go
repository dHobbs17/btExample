package main

import (
	"btexample/tui"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	CreateMainMenu()
}

func CreateMainMenu() {
	p := tea.NewProgram(tui.NewProject())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
