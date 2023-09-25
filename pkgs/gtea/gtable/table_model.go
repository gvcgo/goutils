package gtable

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type TableModel struct {
	table *table.Model
}

func (that *TableModel) Init() tea.Cmd { return nil }

func (that *TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "tab":
			if that.table.Focused() {
				that.table.Blur()
			} else {
				that.table.Focus()
			}
		case "q", "ctrl+c":
			return that, tea.Quit
		case "enter":
			return that, tea.Batch(
				tea.Printf("Let's go to %s!", that.table.SelectedRow()[1]),
			)
		}
	}
	var tModel table.Model
	tModel, cmd = that.table.Update(msg)
	that.table = &tModel
	return that, cmd
}

func (that *TableModel) View() string {
	return baseStyle.Render(that.table.View()) + "\n"
}
