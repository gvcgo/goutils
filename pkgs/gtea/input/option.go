package input

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
)

type Option struct {
	Program *tea.Program
	model   *OptionModel
}

func NewOption(values []string, opts ...TOption) (option *Option) {
	model := NewOptionModel(values, opts...)
	// model.InputModel.Focus()
	option = &Option{
		model:   model,
		Program: tea.NewProgram(model),
	}
	return
}

func (that *Option) Run() {
	if _, err := that.Program.Run(); err != nil {
		gprint.PrintError("%+v", err)
	}
}

func (that *Option) Value() string {
	return that.model.InputModel.Value()
}
