package tui

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    tui_models "yogoart/tui/models"
)

func RunTui() {
    p := tea.NewProgram(tui_models.NewMainMenu())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
