package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type AnimMenu struct {
    choices []string
    cursor int

    rainAnimMenu tea.Model
    //sinusAnimMenu tea.Model
    //bounceAnimMenu tea.Model
    backMenu tea.Model
}

func NewAnimMenu(parentMenu tea.Model) AnimMenu {
    menu := AnimMenu {
        choices: []string { 
            "Rain Animation",
            //"Sinus Wave Animation",
            //"Bouncing Balls Animation",
            "Back" },
        cursor: 0,
        backMenu: parentMenu,
    }

    menu.rainAnimMenu = NewRainAnimMenu(menu)
    return menu
}

func (menu AnimMenu) Init() tea.Cmd {
    return tea.SetWindowTitle("Animations Menu")
}

func (menu AnimMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
            if strings.Contains(choice, "Rain") {
                return menu.rainAnimMenu, nil
            }
            if strings.Contains(choice, "Back") {
                return menu.backMenu, nil
            }
        }
    }

    return menu, nil
}

func (menu AnimMenu) View() string {
    // The header
    s := "Please select one of the follwoing choices:\n\n"

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
