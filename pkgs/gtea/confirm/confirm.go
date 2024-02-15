package confirm

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gvcgo/goutils/pkgs/gtea/program"
)

type COption func(confirm *Confirm)

func WithTitle(title string) COption {
	return func(confirm *Confirm) {
		confirm.model.Title = title
	}
}

type Confirm struct {
	Program *tea.Program
	model   *ConfirmModel
}

func NewConfirm(opts ...COption) (cfm *Confirm) {
	model := &ConfirmModel{}
	cfm = &Confirm{
		Program: tea.NewProgram(model),
		model:   model,
	}
	for _, opt := range opts {
		opt(cfm)
	}
	return
}

func (cf *Confirm) SetProgramOpts(opts ...tea.ProgramOption) {
	if cf.model != nil {
		cf.Program = tea.NewProgram(cf.model, opts...)
	}
}

func (that *Confirm) Run() {
	program.Run(that.Program)
}

func (that *Confirm) Result() bool {
	return that.model.Result()
}
