package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
)

// تعریف keyMap با متدهای مورد نیاز
type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Back  key.Binding
	Help  key.Binding
	Quit  key.Binding
}

// پیاده‌سازی interface مورد نیاز برای help
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},      // ستون اول
		{k.Enter, k.Back},   // ستون دوم
		{k.Help, k.Quit},    // ستون سوم
	}
}

// تعریف کلیدها با توضیحات
var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter", "select / confirm"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc", "b"),
		key.WithHelp("esc/b", "back to menu"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "show help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

type model struct {
	state    string
	cursor   int
	choices  []string
	help     help.Model
	showHelp bool
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
		help:     help.New(),
		showHelp: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// مهم: تنظیم عرض برای help
		m.help.SetWidth(msg.Width)
		
	case tea.KeyPressMsg:
		// اولین اولویت: بررسی کلید help
		switch msg.String() {
		case "?":
			m.showHelp = !m.showHelp
			return m, nil
		}
		
		// اگر help نمایش داده شده، با هر کلید دیگری آن را ببند
		if m.showHelp {
			m.showHelp = false
			return m, nil
		}
		
		// پردازش منوی اصلی
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
			case "enter", " ":
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
		} else {
			// صفحات دیگر
			switch msg.String() {
			case "esc", "b":
				m.state = "menu"
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}
	}
	
	return m, nil
}

func (m model) View() tea.View {
	var (
		normalStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFF"))
		
		selectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff7f6a")).
				Bold(true)
		
		titleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#5fcca3")).
				Align(lipgloss.Center)
		
		subtitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("244")).
				Italic(true).
				Align(lipgloss.Center)
		
		helpStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("240")).
				PaddingTop(1)
	)
	
	// اگر help فعال باشد، صفحه help کامل را نشان بده
	if m.showHelp {
		// مهم: اطمینان از اینکه help.Model درست مقداردهی شده
		helpView := m.help.View(keys)
		fullHelp := helpStyle.Render(helpView)
		return tea.NewView(fullHelp + "\n\nPress any key to close help\n")
	}
	
	// نمایش منوی اصلی
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
		
		// نمایش help کوتاه در پایین
		s += "\n" + helpStyle.Render(m.help.ShortHelpView(keys.ShortHelp()))
		s += "\nPress ? for full help\n"
		
		return tea.NewView(s)
	}
	
	// صفحه ALL ROOMS
	if m.state == "all" {
		s := titleStyle.Render("📌 ALL ROOMS") + "\n\n"
		s += "1. Room 101 (Available)\n"
		s += "2. Room 102 (Booked)\n"
		s += "3. Room 103 (Available)\n"
		s += "4. Room 104 (Booked)\n"
		s += "\n─────────────────────────────\n"
		
		// help مخصوص این صفحه
		pageKeys := []key.Binding{keys.Back, keys.Quit, keys.Help}
		s += m.help.ShortHelpView(pageKeys)
		s += "\n"
		
		return tea.NewView(s)
	}
	
	// صفحه AVAILABLE ROOMS
	if m.state == "available" {
		s := titleStyle.Render("✅ AVAILABLE ROOMS") + "\n\n"
		s += "• Room 101\n"
		s += "• Room 103\n"
		s += "• Room 105\n"
		s += "\n─────────────────────────────\n"
		
		pageKeys := []key.Binding{keys.Back, keys.Quit, keys.Help}
		s += m.help.ShortHelpView(pageKeys)
		s += "\n"
		
		return tea.NewView(s)
	}
	
	// صفحه BOOKED ROOMS
	if m.state == "booked" {
		s := titleStyle.Render("🔴 BOOKED ROOMS") + "\n\n"
		s += "• Room 102 (John Doe)\n"
		s += "• Room 104 (Jane Smith)\n"
		s += "• Room 106 (Bob Johnson)\n"
		s += "\n─────────────────────────────\n"
		
		pageKeys := []key.Binding{keys.Back, keys.Quit, keys.Help}
		s += m.help.ShortHelpView(pageKeys)
		s += "\n"
		
		return tea.NewView(s)
	}
	
	// صفحه ADD ROOM
	if m.state == "add" {
		s := titleStyle.Render("➕ ADD NEW ROOM") + "\n\n"
		s += "Room Number: ____\n"
		s += "Status: [Available] Booked\n"
		s += "\n[DEMO MODE - Form will be here]\n"
		s += "\n─────────────────────────────\n"
		
		pageKeys := []key.Binding{keys.Back, keys.Quit, keys.Help}
		s += m.help.ShortHelpView(pageKeys)
		s += "\n"
		
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