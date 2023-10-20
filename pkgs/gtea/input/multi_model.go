package input

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	mhelpStyle          = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type Inputer interface {
	View() string
}

type InputList struct {
	inputList []Inputer
	nameList  []string
}

func NewInputList() (ipl *InputList) {
	ipl = &InputList{
		inputList: []Inputer{},
		nameList:  []string{},
	}
	return
}

func (that *InputList) Add(name string, ipt Inputer) {
	that.inputList = append(that.inputList, ipt)
	that.nameList = append(that.nameList, name)
}

func (that *InputList) Len() int {
	return len(that.inputList)
}

func (that *InputList) GetByIndex(index int) Inputer {
	return that.inputList[index]
}

func (that *InputList) GetNameByIndex(index int) string {
	return that.nameList[index]
}

func (that *InputList) SetPromptStyleByIndex(index int, style lipgloss.Style) {
	m := that.GetByIndex(index)
	switch mType := m.(type) {
	case textinput.Model:
		mType.PromptStyle = style
	case *OptionModel:
		mType.InputModel.textInput.PromptStyle = style
	default:
	}
}

func (that *InputList) SetTextStyle(index int, style lipgloss.Style) {
	m := that.GetByIndex(index)
	switch mType := m.(type) {
	case textinput.Model:
		mType.TextStyle = style
	case *OptionModel:
		mType.InputModel.textInput.TextStyle = style
	default:
	}
}

func (that *InputList) SetCursorMode(index int, mode cursor.Mode) (cmd tea.Cmd) {
	m := that.GetByIndex(index)
	switch mType := m.(type) {
	case textinput.Model:
		cmd = mType.Cursor.SetMode(mode)
	case *OptionModel:
		cmd = mType.InputModel.textInput.Cursor.SetMode(mode)
	default:
	}
	return
}

func (that *InputList) Focus(index int) (cmd tea.Cmd) {
	m := that.GetByIndex(index)
	switch mType := m.(type) {
	case textinput.Model:
		cmd = mType.Focus()
	case *OptionModel:
		cmd = mType.InputModel.textInput.Focus()
	default:
	}
	return
}

func (that *InputList) Blur(index int) {
	m := that.GetByIndex(index)
	switch mType := m.(type) {
	case textinput.Model:
		mType.Blur()
	case *OptionModel:
		mType.InputModel.textInput.Blur()
	default:
	}
}

func (that *InputList) Value(index int) (v string) {
	m := that.GetByIndex(index)
	switch mType := m.(type) {
	case textinput.Model:
		v = mType.Value()
	case *OptionModel:
		v = mType.InputModel.textInput.Value()
	default:
	}
	return
}

func (that *InputList) IsFocusedOnOption(focusIdx int) (r bool) {
	m := that.GetByIndex(focusIdx)
	switch m.(type) {
	case textinput.Model:
		return false
	case *OptionModel:
		return true
	default:
	}
	return
}

func (that *InputList) Update(index int, msg tea.Msg) (m Inputer, cmd tea.Cmd) {
	m = that.GetByIndex(index)
	switch mType := m.(type) {
	case textinput.Model:
		m, cmd = mType.Update(msg)

	case *OptionModel:
		m, cmd = mType.Update(msg)
	default:
	}
	return
}

type MOption func(ipt *textinput.Model)

func MWithPlaceholder(pHolder string) MOption {
	return func(ipt *textinput.Model) {
		ipt.Placeholder = pHolder
	}
}

func MWithCharlimit(cLimit int) MOption {
	return func(ipt *textinput.Model) {
		ipt.CharLimit = cLimit
	}
}

func MWithWidth(width int) MOption {
	return func(ipt *textinput.Model) {
		ipt.Width = width
	}
}

func MWithEchoChar(echoChar string) MOption {
	runeList := []rune(echoChar)
	return func(ipt *textinput.Model) {
		ipt.EchoCharacter = runeList[0]
	}
}

func MWithEchoMode(echoMode textinput.EchoMode) MOption {
	return func(ipt *textinput.Model) {
		ipt.EchoMode = echoMode
	}
}

type InputMultiModel struct {
	focusIndex int
	inputs     *InputList
	cursorMode cursor.Mode
	submitCmd  tea.Cmd
}

func NewInputMultiModel() (imm *InputMultiModel) {
	imm = &InputMultiModel{
		inputs:    NewInputList(),
		submitCmd: tea.Quit,
	}
	return
}

func (that *InputMultiModel) SetSubmitCmd(scmd tea.Cmd) {
	that.submitCmd = scmd
}

func (that *InputMultiModel) AddOneInput(key string, opts ...MOption) {
	ipt := textinput.New()
	for _, opt := range opts {
		opt(&ipt)
	}
	MWithPlaceholder(key)(&ipt)
	ipt.Cursor.Style = cursorStyle
	that.inputs.Add(key, ipt)
}

func (that *InputMultiModel) AddOneOption(values []string, opts ...MOption) {
	if len(values) == 0 {
		gprint.PrintError("option value list is empty")
		return
	}
	option := NewOptionModel(values)
	for _, opt := range opts {
		opt(option.InputModel.textInput)
	}
	option.InputModel.textInput.Cursor.Style = cursorStyle
	that.inputs.Add(values[0], option)
}

func (that *InputMultiModel) Init() tea.Cmd {
	return textinput.Blink
}

func (that *InputMultiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return that, tea.Quit
		// Change cursor mode
		case "ctrl+r":
			that.cursorMode++
			if that.cursorMode > cursor.CursorHide {
				that.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, that.inputs.Len())
			for i := 0; i < that.inputs.Len(); i++ {
				cmds[i] = that.inputs.SetCursorMode(i, that.cursorMode)
			}
			return that, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && that.focusIndex == that.inputs.Len() {
				return that, that.submitCmd
			}

			if that.inputs.IsFocusedOnOption(that.focusIndex) && (s == "down" || s == "up") {
				_, cmd := that.inputs.Update(that.focusIndex, msg)
				return that, cmd
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				that.focusIndex--
			} else {
				that.focusIndex++
			}

			if that.focusIndex > that.inputs.Len() {
				that.focusIndex = 0
			} else if that.focusIndex < 0 {
				that.focusIndex = that.inputs.Len()
			}

			cmds := make([]tea.Cmd, that.inputs.Len())
			for i := 0; i < that.inputs.Len(); i++ {
				if i == that.focusIndex {
					// Set focused state
					cmds[i] = that.inputs.Focus(i)
					that.inputs.SetPromptStyleByIndex(i, focusedStyle)
					that.inputs.SetTextStyle(i, focusedStyle)
					continue
				}
				// Remove focused state
				that.inputs.Blur(i)
				that.inputs.SetPromptStyleByIndex(i, noStyle)
				that.inputs.SetTextStyle(i, noStyle)
			}

			return that, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := that.updateInputs(msg)

	return that, cmd
}

func (that *InputMultiModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, that.inputs.Len())

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range that.inputs.inputList {
		that.inputs.inputList[i], cmds[i] = that.inputs.Update(i, msg)
	}
	return tea.Batch(cmds...)
}

func (that *InputMultiModel) View() string {
	var b strings.Builder

	for i := range that.inputs.inputList {
		b.WriteString(that.inputs.inputList[i].View())
		if i < that.inputs.Len()-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if that.focusIndex == that.inputs.Len() {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(mhelpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(that.cursorMode.String()))
	b.WriteString(mhelpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

func (that *InputMultiModel) Values() map[string]string {
	r := make(map[string]string)
	for idx, name := range that.inputs.nameList {
		r[name] = that.inputs.Value(idx)
	}
	return r
}
