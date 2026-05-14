package components

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
)

// KeyMap تعریف کلیدهای میانبر
type KeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Back  key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Enter, k.Back, k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Enter, k.Back},
		{k.Help, k.Quit},
	}
}

func NewKeyMap() KeyMap {
	return KeyMap{
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
			key.WithHelp("enter", "select"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc", "b"),
			key.WithHelp("esc/b", "back to menu"),
		),
		Help: key.NewBinding(
			key.WithKeys("h"),
			key.WithHelp("h", "toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

type Help struct {
	Model    help.Model
	ShowHelp bool
	Keys     KeyMap
}


func NewHelp() Help {
	return Help{
		Model:    help.New(),
		ShowHelp: false,
		Keys:     NewKeyMap(),
	}
}

func (h *Help) SetWidth(width int) {
	h.Model.SetWidth(width)
}

func (h *Help) Toggle() {
	h.ShowHelp = !h.ShowHelp
}

func (h *Help) Hide() {
	h.ShowHelp = false
}

func (h Help) IsVisible() bool {
	return h.ShowHelp
}

func (h Help) RenderFull() string {
	return h.Model.FullHelpView(h.Keys.FullHelp())
}

func (h Help) RenderShort() string {
	return h.Model.ShortHelpView(h.Keys.ShortHelp())
}