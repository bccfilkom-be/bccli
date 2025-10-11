package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			switch m.step {
			case stepPackage:
				if m.projectName.Value() == "" {
					break
				}
				m.step = stepFramework
				m.cursor = 0
			case stepFramework:
				m.selectedFramework = m.cursor
				m.step = stepDriver
				m.cursor = 0
			case stepDriver:
				m.selectedDriver = m.cursor
				m.step = stepFinish
				m.cursor = 0
			case stepFinish:
				return m, tea.Quit
			}
		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown:
			switch m.step {
			case stepFramework:
				if m.cursor < len(m.frameworkChoices)-1 {
					m.cursor++
				}
			case stepDriver:
				if m.cursor < len(m.driverChoices)-1 {
					m.cursor++
				}
			case stepFinish:
				if m.cursor < 1 {
					m.cursor++
				}
			}
		}
	}
	var cmd tea.Cmd
	if m.step == stepPackage {
		m.projectName, cmd = m.projectName.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	switch m.step {
	case stepPackage:
		return fmt.Sprintf(
			"Project name?\n%s\n\n(esc or ctrl+c to quit)\n",
			m.projectName.View(),
		)
	case stepFramework:
		return renderList("Choose HTTP Framework", m.frameworkChoices, m.cursor)
	case stepDriver:
		return renderList("Choose Database Driver", m.driverChoices, m.cursor)
	case stepFinish:
		return renderConfirmation(m)
	}
	return ""
}

func renderList(title string, choices []string, cursor int) string {
	s := fmt.Sprintf("%s:\n\n", title)
	for i, choice := range choices {
		prefix := " "
		if cursor == i {
			prefix = ">"
		}
		s += fmt.Sprintf("%s %s\n", prefix, choice)
	}
	s += "\n(↑/↓ to move, Enter to select, Esc/Ctrl+C to quit)\n"
	return s
}

func renderConfirmation(m model) string {
	confirmOptions := []string{"✓ Create Project", "✗ Cancel"}
	s := "\nConfirm your settings:\n\n"
	s += fmt.Sprintf("  Project Name: %s\n", m.projectName.Value())
	s += fmt.Sprintf("  Framework: %s\n", m.frameworkChoices[m.selectedFramework])
	s += fmt.Sprintf("  Driver: %s\n\n", m.driverChoices[m.selectedDriver])

	for i, option := range confirmOptions {
		prefix := " "
		if m.cursor == i {
			prefix = ">"
		}
		s += fmt.Sprintf("%s %s\n", prefix, option)
	}
	s += "\n(↑/↓ to move, Enter to confirm, Esc/Ctrl+C to quit)\n"
	return s
}
