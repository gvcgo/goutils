package gtable

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
)

type Table struct {
	Program *tea.Program
	model   *TableModel
}

func NewTable(opts ...table.Option) (t *Table) {
	tt := table.New(opts...)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	// s.Cell = s.Cell.Foreground(lipgloss.Color("#808080"))
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	tt.SetStyles(s)

	model := &TableModel{table: &tt}
	t = &Table{
		Program: tea.NewProgram(model),
		model:   model,
	}
	return
}

func (that *Table) Run() {
	if _, err := that.Program.Run(); err != nil {
		gprint.PrintError("%+v", err)
	}
}
