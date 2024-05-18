package models

import (
	"strconv"
    "strings"

	rpm "yogoart/models/rain_anim_params"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

var (
    textInputPrompts = []string{
        "Number of Droplet Sources (0-128)",
        "Screen Height (0-32)",
        "Minimum Droplet Velocity (1-99)",
        "Maximum Droplet Velocity (1-99)",
        "Minimum Droplet Length (1-10)",
        "Maximum Droplet Length (1-10)",
        "Swap Axes (true/false)",
    }
)

type RainAnimMenu struct {
    focusIndex int
    inputs []textinput.Model
    cursorMode cursor.Mode

    backMenu tea.Model

    rainParams rpm.RainAnimParams
    err error
}

func NewRainAnimMenu(parentMenu tea.Model) RainAnimMenu {
    menu := RainAnimMenu{
        rainParams: *rpm.NewRainAnimParamsDefaults(),
        inputs: make([]textinput.Model, 7),
        backMenu: parentMenu,
    }

    var t textinput.Model
    for i := range menu.inputs {
        t = textinput.New()
        t.Width = 128

        switch i {
        case 0: // Number of Droplets
            t.Placeholder = strconv.Itoa(int(menu.rainParams.NumDropletSources))
            t.CharLimit = 3
            t.Focus()
        case 1:
            t.Placeholder = strconv.Itoa(int(menu.rainParams.RainHeight))
            t.CharLimit = 3
        case 2:
            t.Placeholder = strconv.Itoa(int(menu.rainParams.MinVelocity))
            t.CharLimit = 2
        case 3:
            t.Placeholder = strconv.Itoa(int(menu.rainParams.MaxVelocity))
            t.CharLimit = 2
        case 4:
            t.Placeholder = strconv.Itoa(int(menu.rainParams.DropletMinLen))
            t.CharLimit = 2
        case 5:
            t.Placeholder = strconv.Itoa(int(menu.rainParams.DropletMaxLen))
            t.CharLimit = 2
        case 6:
            t.Placeholder = strconv.FormatBool(menu.rainParams.IsSwapAxes)
            t.CharLimit = 5
        }

        menu.inputs[i] = t
    }

    return menu
}

func (menu RainAnimMenu) Init() tea.Cmd {
    return textinput.Blink
}

func (menu RainAnimMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return menu, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			menu.cursorMode++
			if menu.cursorMode > cursor.CursorHide {
				menu.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(menu.inputs))
			for i := range menu.inputs {
				cmds[i] = menu.inputs[i].Cursor.SetMode(menu.cursorMode)
			}
			return menu, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && menu.focusIndex == len(menu.inputs) {
				return menu.backMenu, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				menu.focusIndex--
			} else {
				menu.focusIndex++
			}

			if menu.focusIndex > len(menu.inputs) {
				menu.focusIndex = 0
			} else if menu.focusIndex < 0 {
				menu.focusIndex = len(menu.inputs)
			}

			cmds := make([]tea.Cmd, len(menu.inputs))
			for i := 0; i <= len(menu.inputs)-1; i++ {
				if i == menu.focusIndex {
					// Set focused state
					cmds[i] = menu.inputs[i].Focus()
					continue
				}
				// Remove focused state
				menu.inputs[i].Blur()
			}

			return menu, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := menu.updateInputs(msg)
	return menu, cmd
}

func (menu *RainAnimMenu) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(menu.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range menu.inputs {
		menu.inputs[i], cmds[i] = menu.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (menu RainAnimMenu) View() string {
	var b strings.Builder

	for i := range menu.inputs {
        b.WriteString(textInputPrompts[i] + ": " + menu.inputs[i].View())
		if i < len(menu.inputs)-1 {
			b.WriteRune('\n')
		}
	}

    return b.String()
}

