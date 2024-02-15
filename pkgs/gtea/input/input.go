package input

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gvcgo/goutils/pkgs/gtea/program"
)

type Input struct {
	Program *tea.Program
	model   *InputModel
}

func NewInput(opts ...TOption) (ipt *Input) {
	model := NewInputModel(opts...)
	model.Focus()
	for _, opt := range opts {
		opt(model)
	}
	ipt = &Input{
		model:   model,
		Program: tea.NewProgram(model),
	}
	return
}

func (that *Input) SetProgramOpts(opts ...tea.ProgramOption) {
	if that.model != nil {
		that.Program = tea.NewProgram(that.model, opts...)
	}
}

func (that *Input) Run() {
	program.Run(that.Program)
}

func (that *Input) Value() string {
	return that.model.Value()
}
