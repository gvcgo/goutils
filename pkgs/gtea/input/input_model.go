package input

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TOption func(ipm *InputModel)

func WithPlaceholder(pHolder string) TOption {
	return func(ipm *InputModel) {
		ipm.textInput.Placeholder = pHolder
	}
}

func WithCharlimit(cLimit int) TOption {
	return func(ipm *InputModel) {
		ipm.textInput.CharLimit = cLimit
	}
}

func WithWidth(width int) TOption {
	return func(ipm *InputModel) {
		ipm.textInput.Width = width
	}
}

func WithEchoChar(echoChar string) TOption {
	runeList := []rune(echoChar)
	return func(ipm *InputModel) {
		ipm.textInput.EchoCharacter = runeList[0]
	}
}

func WithEchoMode(echoMode textinput.EchoMode) TOption {
	return func(ipm *InputModel) {
		ipm.textInput.EchoMode = echoMode
	}
}

var (
	helpStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
	focusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

type ErrMsg error

type InputModel struct {
	textInput *textinput.Model
	err       error
}

func NewInputModel(opts ...TOption) (im *InputModel) {
	ti := textinput.New()
	ti.Cursor.Style = focusStyle
	ti.PromptStyle = focusStyle
	ti.TextStyle = focusStyle
	im = &InputModel{
		textInput: &ti,
	}

	for _, opt := range opts {
		opt(im)
	}
	return
}

func (that *InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (that *InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
		ti  textinput.Model
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return that, tea.Quit
		}

	// We handle errors just like any other message
	case ErrMsg:
		that.err = msg
		return that, nil
	}

	ti, cmd = that.textInput.Update(msg)
	that.textInput = &ti
	return that, cmd
}

func (that *InputModel) View() string {
	return fmt.Sprintf(
		"%s\n%s\n",
		that.textInput.View(),
		helpStyle(`press "esc" to quit`),
	) + "\n"
}

func (that *InputModel) Value() string {
	return that.textInput.Value()
}

func (that *InputModel) Focus() {
	that.textInput.Focus()
}
