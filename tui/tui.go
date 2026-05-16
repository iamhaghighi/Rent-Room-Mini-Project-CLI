package main

import (
	"fmt"
	"os"
	"rent_room/components"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF"))

	// cursor
	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff7f6a")).
			Bold(true)
	// #ff7f6a - #7fd6b3

	// title
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5fcca3")).
			Align(lipgloss.Center)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("244")).
			Italic(true).
			Align(lipgloss.Center)
)

type model struct {
	state   string
	cursor  int
	choices []string

	table components.TableModel
	help  components.Help
}

func initialModel() model {
	return model{
		state:  "menu",
		cursor: 0,
		choices: []string{
			"All Rooms",
			"Available Rooms",
			"Booked Rooms",
			"Add Room",
			"Exit",
		},
		table: components.NewStyleTable(),
		help: components.NewHelp(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.help.SetWidth(msg.Width)

	case tea.KeyPressMsg:

		//* help
		switch msg.String() {
		case "h":
			m.help.Toggle()
			return m, nil
		}
		if m.help.ShowHelp {
			switch msg.String() {
			case "esc", "b":
				m.help.Toggle()
				return m, nil
			}
		}

		if m.state == "menu" {

			switch msg.String() {

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			case "enter", "space":
				switch m.cursor {
				case 0:
					m.state = "all"
				case 1:
					m.state = "available"
				case 2:
					m.state = "booked"
				case 3:
					m.state = "add"
				case 4:
					return m, tea.Quit
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}

		if m.state != "menu" {
			switch msg.String() {
			case "esc", "b":
				m.state = "menu"
			}
		}
	}

	return m, nil
}

func (m model) View() tea.View {

	if m.help.ShowHelp {
		helpView := m.help.Model.FullHelpView(m.help.Keys.FullHelp())
		return tea.NewView(helpView + "\n\nPress 'h' to back to menu\n")
	}

	if m.state == "menu" {
		title := titleStyle.Render("🏨 Hotel Rental System")
		subtitle := subtitleStyle.Render("Manage Room Easy")

		width := lipgloss.Width(title) + 4
		if width < 34 {
			width = 34
		}

		top := "╭" + strings.Repeat("─", width-2) + "╮"
		bottom := "╰" + strings.Repeat("─", width-2) + "╯"

		s := "\n" + top + "\n"
		s += "│" + lipgloss.Place(width-2, 1, lipgloss.Center, lipgloss.Center, title) + "│\n"
		s += "│" + lipgloss.Place(width-2, 1, lipgloss.Center, lipgloss.Center, subtitle) + "│\n"
		s += bottom + "\n\n"

		for i, choice := range m.choices {
			cursor := "[ ]"
			itemStyle := normalStyle

			if m.cursor == i {
				cursor = selectedStyle.Render("[>]")
				itemStyle = selectedStyle
			}

			s += fmt.Sprintf("%s %s\n", cursor, itemStyle.Render(choice))
		}

		s += "\n" + m.help.Model.ShortHelpView(m.help.Keys.ShortHelp())

		return tea.NewView(s)
	}

	if m.state == "all" {
		return m.table.RenderViewTable()
	}

	if m.state == "available" {
		s := "\n✅ AVAILABLE ROOMS\n\n"
		s += "• Room 101\n"
		s += "• Room 103\n"
		s += "• Room 105\n"
		s += "\n─────────────────────────────"
		s += "\nPress 'esc' or 'b' to back to menu\n"
		return tea.NewView(s)
	}

	if m.state == "booked" {
		s := "\n🔴 BOOKED ROOMS\n\n"
		s += "• Room 102 (John Doe)\n"
		s += "• Room 104 (Jane Smith)\n"
		s += "• Room 106 (Bob Johnson)\n"
		s += "\n─────────────────────────────"
		s += "\nPress 'esc' or 'b' to back to menu\n"
		return tea.NewView(s)
	}

	if m.state == "add" {
		s := "\n➕ ADD NEW ROOM\n\n"
		s += "Room Number: ____\n"
		s += "Status: [Available] Booked\n"
		s += "\n[DEMO MODE - Press any key to back]\n"
		s += "\n─────────────────────────────"
		s += "\nPress 'esc' or 'b' to back to menu\n"
		return tea.NewView(s)
	}

	return tea.NewView("Unknown state")
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
