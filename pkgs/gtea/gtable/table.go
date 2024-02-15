package gtable

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gvcgo/goutils/pkgs/gtea/program"
)

type Table struct {
	Program *tea.Program
	model   *TableModel
}

func NewTable(opts ...Option) (t *Table) {
	tt := New(opts...)
	s := DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Cell = s.Cell.Align(lipgloss.Left)
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

func (that *Table) SetProgramOpts(opts ...tea.ProgramOption) {
	if that.model != nil {
		that.Program = tea.NewProgram(that.model, opts...)
	}
}

func (that *Table) Run() {
	program.Run(that.Program)
}
