package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	JumpToString key.Binding
	Next         key.Binding
	Previous     key.Binding
	Left         key.Binding
	Right        key.Binding
	Select       key.Binding
	Mute         key.Binding
	Quit         key.Binding
	Help         key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Select, k.Mute, k.Quit, k.Help,
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.JumpToString, k.Previous, k.Next, k.Left, k.Right},
		{k.Select, k.Mute, k.Quit, k.Help},
	}
}

var keys = keyMap{
	JumpToString: key.NewBinding(
		key.WithKeys("1", "2", "3", "4", "5", "6", "7", "8", "9"),
		key.WithHelp("0..9", "jump to string 0..9"),
	),
	Next: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "next"),
	),
	Previous: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "previous"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "right"),
	),
	Select: key.NewBinding(
		key.WithKeys(" ", "enter"),
		key.WithHelp("space/enter", "select"),
	),
	Mute: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "mute"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
}
