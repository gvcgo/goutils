package input

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
)

type OptionModel struct {
	*InputModel
	ValueList []string
	Idx       int
}

func NewOptionModel(values []string, opts ...TOption) (om *OptionModel) {
	if len(values) == 0 {
		gprint.PrintError("option values cannot be empty")
		return
	}
	im := NewInputModel(opts...)
	om = &OptionModel{
		InputModel: im,
		ValueList:  values,
	}
	om.InputModel.textInput.SetValue(values[om.Idx])
	return
}

func (that *OptionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "enter":
			return that, tea.Quit
		case "up":
			if that.Idx < len(that.ValueList)-1 {
				that.Idx++
			} else {
				that.Idx = 0
			}
			that.InputModel.textInput.SetValue(that.ValueList[that.Idx])
		case "down":
			if that.Idx > 0 {
				that.Idx--
			} else {
				that.Idx = len(that.ValueList) - 1
			}
			that.InputModel.textInput.SetValue(that.ValueList[that.Idx])
		default:
		}
	// We handle errors just like any other message
	case ErrMsg:
		that.err = msg
		return that, nil
	}
	return that, nil
}

func (that *OptionModel) View() string {
	return fmt.Sprintf(
		"%s\n",
		that.textInput.View(),
	)
}

func (that *OptionModel) SetPromptStyle(style lipgloss.Style) {
	that.InputModel.textInput.PromptStyle = style
}

func (that *OptionModel) SetTextStyle(style lipgloss.Style) {
	that.InputModel.textInput.TextStyle = style
}

func (that *OptionModel) Focus() tea.Cmd {
	return that.InputModel.textInput.Focus()
}

func (that *OptionModel) Value() string {
	return that.InputModel.textInput.Value()
}

func (that *OptionModel) Blur() {
	that.InputModel.textInput.Blur()
}

func (that *OptionModel) SetCursorMode(mode cursor.Mode) tea.Cmd {
	return that.InputModel.textInput.Cursor.SetMode(mode)
}

func (that *OptionModel) IsOption() bool {
	return true
}

func (that *OptionModel) SetValue(v string) {
	that.InputModel.textInput.SetValue(v)
}
