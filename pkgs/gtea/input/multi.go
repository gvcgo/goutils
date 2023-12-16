package input

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
	"github.com/moqsien/goutils/pkgs/gtea/program"
)

type MultiInput struct {
	Program *tea.Program
	model   *InputMultiModel
}

func NewMultiInput() (mipt *MultiInput) {
	m := NewInputMultiModel()
	mipt = &MultiInput{
		model: m,
	}
	return
}

func (that *MultiInput) SetProgramOpts(opts ...tea.ProgramOption) {
	if that.model != nil {
		that.Program = tea.NewProgram(that.model, opts...)
	}
}

func (that *MultiInput) AddOneItem(key string, opts ...MOption) {
	that.model.AddOneInput(key, opts...)
}

func (that *MultiInput) AddOneOption(name string, values []string, opts ...MOption) {
	that.model.AddOneOption(name, values, opts...)
}

func (that *MultiInput) Run() {
	if that.model.inputs.Len() == 0 {
		gprint.PrintError("No item is added!")
		return
	}
	if that.Program == nil {
		that.Program = tea.NewProgram(that.model)
	}
	program.Run(that.Program)
}

func (that *MultiInput) Values() map[string]string {
	return that.model.Values()
}
