package input

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
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

func WithPlaceholderStyle(style lipgloss.Style) TOption {
	return func(ipm *InputModel) {
		ipm.textInput.PlaceholderStyle = style
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

func WithPrompt(prompt string) TOption {
	return func(ipm *InputModel) {
		if !strings.HasSuffix(prompt, ": ") {
			prompt += ": "
		}
		ipm.textInput.Prompt = prompt
	}
}

var (
	helpStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Render
	focusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

type ErrMsg error

type InputModel struct {
	textInput *textinput.Model
	err       error
	helpStr   string
}

func NewInputModel(opts ...TOption) (im *InputModel) {
	ti := textinput.New()
	ti.Cursor.Style = focusStyle
	ti.TextStyle = focusStyle
	ti.PlaceholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#008000"))
	im = &InputModel{
		textInput: &ti,
		helpStr:   helpStyle(`Press "Enter" to end input, "Esc" to quit.`),
	}
	for _, opt := range opts {
		opt(im)
	}
	return
}

func (that *InputModel) SetHelpStr(help string) {
	if help == "" {
		that.helpStr = ""
		return
	}
	that.helpStr = helpStyle(help)
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
	r := fmt.Sprintf("%s\n", that.textInput.View())
	if that.helpStr != "" {
		r = fmt.Sprintf(
			"%s\n%s\n",
			that.textInput.View(),
			helpStyle(`Press "Enter" to end input, "Esc" to quit.`),
		)
	}
	return r
}

func (that *InputModel) SetPromptStyle(style lipgloss.Style) {
	that.textInput.PromptStyle = style
}

func (that *InputModel) SetTextStyle(style lipgloss.Style) {
	that.textInput.TextStyle = style
}

func (that *InputModel) Value() string {
	return that.textInput.Value()
}

func (that *InputModel) Focus() tea.Cmd {
	return that.textInput.Focus()
}

func (that *InputModel) Blur() {
	that.textInput.Blur()
}

func (that *InputModel) SetCursorMode(mode cursor.Mode) tea.Cmd {
	return that.textInput.Cursor.SetMode(mode)
}

func (that *InputModel) IsOption() bool {
	return false
}

func (that *InputModel) SetValue(v string) {
	that.textInput.SetValue(v)
}
