package models

import (
    tea "github.com/charmbracelet/bubbletea"
)

type Menu struct {
    mainMenu tea.Model
    currentMenu tea.Model
}

func NewMenu() Menu {
    menu := Menu{
        mainMenu: NewMainMenu(),
    }
    menu.currentMenu = menu.mainMenu
    return menu
}

func (menu Menu) Init() tea.Cmd {
    return nil
}

func (menu Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    menu.currentMenu, cmd = menu.currentMenu.Update(msg)
    return menu, cmd
}

func (menu Menu) View() string {
    return menu.currentMenu.View()
}
