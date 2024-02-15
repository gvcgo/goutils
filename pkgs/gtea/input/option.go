package input

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gvcgo/goutils/pkgs/gtea/program"
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

func (that *Option) SetProgramOpts(opts ...tea.ProgramOption) {
	if that.model != nil {
		that.Program = tea.NewProgram(that.model, opts...)
	}
}

func (that *Option) Run() {
	program.Run(that.Program)
}

func (that *Option) Value() string {
	return that.model.InputModel.Value()
}
