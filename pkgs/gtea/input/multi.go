package input

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
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

func (that *MultiInput) AddOneItem(key string, opts ...MOption) {
	that.model.AddOneInput(key, opts...)
}

func (that *MultiInput) Run() {
	if that.model.inputs.Len() == 0 {
		gprint.PrintError("No item is added!")
		return
	}
	if that.Program == nil {
		that.Program = tea.NewProgram(that.model)
	}
	if _, err := that.Program.Run(); err != nil {
		gprint.PrintError("%+v", err)
	}
}

func (that *MultiInput) Values() map[string]string {
	return that.model.Values()
}
