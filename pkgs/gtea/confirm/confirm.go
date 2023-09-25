package confirm

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
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

func (that *Confirm) Run() {
	if _, err := that.Program.Run(); err != nil {
		gprint.PrintError("%+v", err)
	}
}

func (that *Confirm) Result() bool {
	return that.model.Result()
}
