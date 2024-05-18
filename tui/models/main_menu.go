package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MainMenu struct {
    choices []string
    cursor int

    animMenu tea.Model
}

func NewMainMenu() MainMenu {
    menu := MainMenu{
        choices: []string { "Animations", "Exit" },
        cursor: 0,
    }
    menu.animMenu = NewAnimMenu(menu)
    return menu
}

func (menu MainMenu) Init() tea.Cmd {
    return tea.SetWindowTitle("Main Menu")
}

func (menu MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return menu, tea.Quit

        case "up", "k":
            if menu.cursor > 0 {
                menu.cursor--
            }

        case "down", "j":
            if menu.cursor < len(menu.choices)-1 {
                menu.cursor++
            }

        case "enter", " ":
            choice := menu.choices[menu.cursor]
            if strings.Contains(choice, "Animations") {
                return menu.animMenu, nil
            }
            if strings.Contains(choice, "Exit") {
                return menu, tea.Quit
            }
        }
    }

    return menu, nil
}

func (menu MainMenu) View() string {
    // The header
    s := "Welcome to YoGoArt\n\n"
    s += "Please select one of the follwoing choices:\n\n"

    for i, choice := range menu.choices {
        cursor := " " // no cursor
        if menu.cursor == i {
            cursor = ">" // cursor!
        }

        // Render the row
        s += fmt.Sprintf("%s %s\n", cursor, choice)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}
