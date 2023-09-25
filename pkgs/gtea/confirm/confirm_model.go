package confirm

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
)

type ConfirmModel struct {
	Title string
	Ok    bool
	key   string
}

func (that *ConfirmModel) Init() tea.Cmd { return nil }

func (that *ConfirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "Y", "y":
			that.Ok = true
			that.key = msg.String()
			return that, tea.Quit
		case "N", "n":
			that.Ok = false
			that.key = msg.String()
			return that, tea.Quit
		}
	}
	return that, nil
}

func (that *ConfirmModel) View() string {
	title := gprint.CyanStr("%s [Y/N].", that.Title)
	if that.key != "" {
		title += fmt.Sprintf(" %s", gprint.YellowStr(that.key))
	}
	return title + "\n"
}

func (that *ConfirmModel) Result() bool {
	return that.Ok
}
