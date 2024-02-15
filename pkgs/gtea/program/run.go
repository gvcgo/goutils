package program

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gvcgo/goutils/pkgs/gtea/gprint"
)

func Run(app *tea.Program) {
	if _, err := app.Run(); err != nil {
		gprint.PrintError("%+v", err)
	}
}
