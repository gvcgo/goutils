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

func (d ItemDelegate) Clear() {
	for key := range *d.chosen {
		delete(*d.chosen, key)
	}
}

func (d ItemDelegate) Delete(key Item) {
	delete(*d.chosen, key)
}

func (d ItemDelegate) Add(key Item) {
	(*d.chosen)[key] = struct{}{}
}

func (d ItemDelegate) GetAll() map[Item]struct{} {
	return *d.chosen
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

func WidthEnableMulti(enable bool) SOption {
	return func(sm *SelectorModel) {
		sm.multi = enable
	}
}

const (
	DefaultWidth  = 20
	DefaultHeight = 10
)

type SelectorModel struct {
	list      *list.Model
	delegate  *ItemDelegate
	items     []Item
	quitting  bool
	multi     bool
	submitCmd tea.Cmd
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
	l.SetShowHelp(false)
	sm = &SelectorModel{
		list:      &l,
		delegate:  delegate,
		items:     items,
		submitCmd: tea.Quit,
	}
	for _, opt := range opts {
		opt(sm)
	}
	return
}

func (that *SelectorModel) SetSubmitCmd(scmd tea.Cmd) {
	that.submitCmd = scmd
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
		case "ctrl+c", "q":
			that.quitting = true
			return that, tea.Quit
		case "tab", "esc":
			cmd := that.submitCmd
			if len(*that.delegate.chosen) == 0 {
				cmd = nil
			}
			return that, cmd
		case "enter":
			i, ok := that.list.SelectedItem().(Item)
			if ok && that.multi {
				if _, exist := (*that.delegate.chosen)[i]; !exist {
					that.delegate.Add(i)
				} else {
					that.delegate.Delete(i)
				}
			} else if ok && !that.multi {
				that.delegate.Clear()
				that.delegate.Add(i)
			}
			cmd := that.submitCmd
			if that.multi {
				cmd = nil
			}
			return that, cmd
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

var helpStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#FFA500")).Render

func (that *SelectorModel) View() string {
	var r string
	if that.multi {
		r = lipgloss.JoinVertical(
			lipgloss.Left,
			that.list.View(),
			helpStyle("↑/k up • ↓/j down • / filter • q quit • ? more"),
			helpStyle("Press Enter to select one item."),
			helpStyle(`Press Tab/Esc to confirm selections.`),
		)
	} else {
		r = lipgloss.JoinVertical(
			lipgloss.Left,
			that.list.View(),
			helpStyle("↑/k up • ↓/j down • / filter • q quit • ? more"),
			helpStyle(`Press Enter to select one item.`),
		)
	}
	return r
}

func (that *SelectorModel) ChosenList() (r []string) {
	for item := range that.delegate.GetAll() {
		r = append(r, string(item))
	}
	// choose the first item by default if none was chosen.
	if len(r) == 0 && len(that.items) > 0 {
		r = append(r, string(that.items[0]))
	}
	return
}

func (that *SelectorModel) Values() (r map[string]string) {
	r = map[string]string{}
	for _, item := range that.ChosenList() {
		r[item] = item
	}
	return
}
