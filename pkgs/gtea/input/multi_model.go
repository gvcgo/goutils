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
	Init() tea.Cmd
	View() string
	Update(tea.Msg) (tea.Model, tea.Cmd)
	Focus() tea.Cmd
	Value() string
	SetPromptStyle(style lipgloss.Style)
	SetTextStyle(style lipgloss.Style)
	Blur()
	SetCursorMode(mode cursor.Mode) tea.Cmd
	IsOption() bool
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

func MWithPrompt(prompt string) MOption {
	return func(ipt *textinput.Model) {
		if !strings.HasSuffix(prompt, ": ") {
			prompt += ": "
		}
		ipt.Prompt = prompt
	}
}

type InputMultiModel struct {
	focusIndex         int
	inputs             *InputList
	cursorMode         cursor.Mode
	submitCmd          tea.Cmd
	inputPromptPattern string
}

func NewInputMultiModel() (imm *InputMultiModel) {
	imm = &InputMultiModel{
		inputs:             NewInputList(),
		submitCmd:          tea.Quit,
		inputPromptPattern: "%-10s",
	}
	return
}

func (that *InputMultiModel) SetInputPromptPattern(pattern string) {
	that.inputPromptPattern = pattern
}

func (that *InputMultiModel) SetSubmitCmd(scmd tea.Cmd) {
	that.submitCmd = scmd
}

func (that *InputMultiModel) AddOneInput(key string, opts ...MOption) {
	ipt := NewInputModel()
	ipt.SetHelpStr("")
	for _, opt := range opts {
		opt(ipt.textInput)
	}
	MWithPlaceholder(key)(ipt.textInput)
	MWithPrompt(fmt.Sprintf(that.inputPromptPattern, key))(ipt.textInput)
	ipt.textInput.Cursor.Style = cursorStyle
	that.inputs.Add(key, ipt)
	if that.inputs.Len() == 1 {
		that.inputs.GetByIndex(0).Focus()
	}
}

func (that *InputMultiModel) AddOneOption(name string, values []string, opts ...MOption) {
	if len(values) == 0 {
		gprint.PrintError("option value list is empty")
		return
	}
	option := NewOptionModel(values, WithPrompt(fmt.Sprintf(that.inputPromptPattern, name)))
	for _, opt := range opts {
		opt(option.InputModel.textInput)
	}
	option.InputModel.textInput.Cursor.Style = cursorStyle
	that.inputs.Add(name, option)
	if that.inputs.Len() == 1 {
		that.inputs.GetByIndex(0).Focus()
	}
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
				cmds[i] = that.inputs.GetByIndex(i).SetCursorMode(that.cursorMode)
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

			if that.focusIndex >= 0 && that.focusIndex < that.inputs.Len() && that.inputs.GetByIndex(that.focusIndex).IsOption() && (s == "down" || s == "up") {
				_, cmd := that.inputs.GetByIndex(that.focusIndex).Update(msg)
				return that, cmd
			}

			// Cycle indexes, including [Submit]
			switch s {
			case "up", "shift+tab":
				if that.focusIndex < 1 {
					that.focusIndex = that.inputs.Len()
				} else {
					that.focusIndex--
				}
			default:
				if that.focusIndex > that.inputs.Len()-1 {
					that.focusIndex = 0
				} else {
					that.focusIndex++
				}
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
					cmds[i] = that.inputs.GetByIndex(i).Focus()
					that.inputs.GetByIndex(i).SetPromptStyle(focusedStyle)
					that.inputs.GetByIndex(i).SetTextStyle(focusedStyle)
					continue
				}
				// Remove focused state
				that.inputs.GetByIndex(i).Blur()
				that.inputs.GetByIndex(i).SetPromptStyle(noStyle)
				that.inputs.GetByIndex(i).SetTextStyle(noStyle)
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
		_, cmds[i] = that.inputs.GetByIndex(i).Update(msg)
	}
	return tea.Batch(cmds...)
}

func (that *InputMultiModel) View() string {
	rows := []string{}
	for i := range that.inputs.inputList {
		rows = append(rows, that.inputs.GetByIndex(i).View())
	}
	button := &blurredButton
	if that.focusIndex == that.inputs.Len() {
		button = &focusedButton
	}
	rows = append(rows, *button)

	helpStr := mhelpStyle.Render("cursor mode is ") + cursorModeHelpStyle.Render(that.cursorMode.String()) + mhelpStyle.Render(" (ctrl+r to change style)")
	rows = append(rows, helpStr)
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func (that *InputMultiModel) Values() map[string]string {
	r := make(map[string]string)
	for idx, name := range that.inputs.nameList {
		r[name] = that.inputs.GetByIndex(idx).Value()
	}
	return r
}
