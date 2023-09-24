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
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4).Foreground(lipgloss.Color("#DCDCDC"))
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#FF00FF"))
)

type ItemDelegate struct {
	chosen *map[Item]struct{}
}

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

	hintStr := "✔ "
	if _, exists := (*d.chosen)[i]; !exists {
		hintStr = "> "
	}
	fn := func(s ...string) string {
		return itemStyle.Render(hintStr, strings.Join(s, " "))
	}
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render(hintStr + strings.Join(s, " "))
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
	DefaultHeight = 10
)

type SelectorModel struct {
	list     *list.Model
	delegate *ItemDelegate
	items    []Item
	quitting bool
}

func NewSelectorModel(items []Item, opts ...SOption) (sm *SelectorModel) {
	itemList := []list.Item{}
	for _, item := range items {
		itemList = append(itemList, item)
	}
	delegate := &ItemDelegate{
		chosen: &map[Item]struct{}{},
	}
	l := list.New(itemList, delegate, DefaultWidth, DefaultHeight)
	sm = &SelectorModel{
		list:     &l,
		delegate: delegate,
		items:    items,
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
		case "esc", "tab", "q":
			return that, tea.Quit
		case "enter":
			i, ok := that.list.SelectedItem().(Item)
			if ok {
				if _, exist := (*that.delegate.chosen)[i]; !exist {
					(*that.delegate.chosen)[i] = struct{}{}
				} else {
					delete((*that.delegate.chosen), i)
				}
			}
			return that, nil
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

func (that *SelectorModel) ChosenList() (r []string) {
	for item := range *(that.delegate.chosen) {
		r = append(r, string(item))
	}
	// choose the first item by default if none was chosen.
	if len(r) == 0 && len(that.items) > 0 {
		r = append(r, string(that.items[0]))
	}
	return
}
