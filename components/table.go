package components

import (
	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))


type TableModel struct {
	table table.Model
}

func defaultTable() table.Model {
	var columns = []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Type", Width: 8},
		{Title: "Bed Count", Width: 11},
		{Title: "Price", Width: 8},
		{Title: "Status", Width: 8},
	}
	var rows = []table.Row{}

	return table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
		table.WithWidth(42),
	)
}

func NewStyleTable() TableModel {
	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t := defaultTable()

	t.SetStyles(s)

	return TableModel{t}

}

func (t TableModel) Focus() {
	if t.table.Focused() {
		t.table.Blur()
	} else {
		t.table.Focus()
	}
}

func (t TableModel) RenderViewTable() tea.View {
	return tea.NewView(baseStyle.Render(t.table.View()) + "\n  " + t.table.HelpView() + "\n")
}
