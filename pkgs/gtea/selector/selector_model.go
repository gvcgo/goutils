package selector

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Item string

func (that Item) FilterValue() string { return "" }

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4).Foreground(lipgloss.Color("#808080"))
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#FF00FF"))
)

type ItemDelegate struct{}

func (d ItemDelegate) Height() int {
	return 1
}

func (d ItemDelegate) Spacing() int {
	return 0
}

func (d ItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (d ItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type SOption func(sm *SelectorModel)

func WithTitle(title string) SOption {
	return func(sm *SelectorModel) {
		sm.list.Title = title
	}
}

func WithShowStatusBar(show bool) SOption {
	return func(sm *SelectorModel) {
		sm.list.SetShowStatusBar(show)
	}
}

func WithFilteringEnabled(show bool) SOption {
	return func(sm *SelectorModel) {
		sm.list.SetFilteringEnabled(show)
	}
}

func WithWidth(width int) SOption {
	return func(sm *SelectorModel) {
		sm.list.SetWidth(width)
	}
}

func WithHeight(height int) SOption {
	return func(sm *SelectorModel) {
		sm.list.SetHeight(height)
	}
}

func WithEnbleInfinite(enable bool) SOption {
	return func(sm *SelectorModel) {
		sm.list.InfiniteScrolling = enable
	}
}

const (
	DefaultWidth  = 20
	DefaultHeight = 14
)

// TODO: support multiple selection and goto start
type SelectorModel struct {
	list     *list.Model
	choice   string
	quitting bool
}

func NewSelectorModel(items []Item, opts ...SOption) (sm *SelectorModel) {
	itemList := []list.Item{}
	for _, item := range items {
		itemList = append(itemList, item)
	}
	l := list.New(itemList, &ItemDelegate{}, DefaultWidth, DefaultHeight)
	sm = &SelectorModel{
		list: &l,
	}
	for _, opt := range opts {
		opt(sm)
	}
	return
}

func (that *SelectorModel) Init() tea.Cmd {
	return nil
}

func (that *SelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		that.list.SetWidth(msg.Width)
		return that, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			that.quitting = true
			return that, tea.Quit

		case "enter":
			i, ok := that.list.SelectedItem().(Item)
			if ok {
				that.choice = string(i)
			}
			return that, tea.Quit
		}
	}

	var (
		cmd   tea.Cmd
		sList list.Model
	)
	sList, cmd = that.list.Update(msg)
	that.list = &sList
	return that, cmd
}

func (that *SelectorModel) View() string {
	return "\n" + that.list.View() + "\n"
}

func (that *SelectorModel) Choice() string {
	return that.choice
}
