package input

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

type InputList struct {
	inputList []textinput.Model
	nameList  []string
}

func NewInputList() (ipl *InputList) {
	ipl = &InputList{
		inputList: []textinput.Model{},
		nameList:  []string{},
	}
	return
}

func (that *InputList) Add(name string, ipt textinput.Model) {
	that.inputList = append(that.inputList, ipt)
	that.nameList = append(that.nameList, name)
}

func (that *InputList) Len() int {
	return len(that.inputList)
}

func (that *InputList) GetByIndex(index int) textinput.Model {
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
				cmds[i] = that.inputs.inputList[i].Cursor.SetMode(that.cursorMode)
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
					cmds[i] = that.inputs.inputList[i].Focus()
					that.inputs.inputList[i].PromptStyle = focusedStyle
					that.inputs.inputList[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				that.inputs.inputList[i].Blur()
				that.inputs.inputList[i].PromptStyle = noStyle
				that.inputs.inputList[i].TextStyle = noStyle
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
	for i, ipt := range that.inputs.inputList {
		that.inputs.inputList[i], cmds[i] = ipt.Update(msg)
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
		r[name] = that.inputs.inputList[idx].Value()
	}
	return r
}
