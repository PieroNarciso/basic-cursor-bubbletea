package cursor

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	posX   int
	posY   int
	height int
	width  int
}

func InitProgram() tea.Model {
	return &model{0, 0, 0, 0}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "j", tea.KeyDown.String():
			m.posY++
		case "k", tea.KeyUp.String():
			m.posY--
		case "h", tea.KeyLeft.String():
			m.posX--
		case "l", tea.KeyRight.String():
			m.posX++
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	}
	return m, nil
}

func (m *model) View() string {
	builder := strings.Builder{}

	for i := 0; i < m.height-1; i++ {
		for j := 0; j < m.width-1; j++ {
			if i == m.posY && j == m.posX {
				builder.WriteRune('*')
			} else {
				builder.WriteRune(' ')
			}
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}
