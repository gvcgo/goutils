package gtea

import (
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type ModelProgressBar struct {
	Progress progress.Model
	err      error
}

func (that *ModelProgressBar) Init() tea.Cmd {
	return nil
}

func (that *ModelProgressBar) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return that, tea.Quit
	case tea.WindowSizeMsg:
		that.Progress.Width = msg.Width - padding*2 - 4
		if that.Progress.Width > maxWidth {
			that.Progress.Width = maxWidth
		}
		return that, nil

	case progressErrMsg:
		that.err = msg.err
		return that, tea.Quit

	case progressMsg:
		var cmds []tea.Cmd

		if msg >= 1.0 {
			cmds = append(cmds, tea.Sequence(finalPause(), tea.Quit))
		}

		cmds = append(cmds, that.Progress.SetPercent(float64(msg)))
		return that, tea.Batch(cmds...)

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := that.Progress.Update(msg)
		that.Progress = progressModel.(progress.Model)
		return that, cmd

	default:
		return that, nil
	}
}

func (that *ModelProgressBar) View() string {
	if that.err != nil {
		return "Error downloading: " + that.err.Error() + "\n"
	}

	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + that.Progress.View() + "\n\n" +
		pad + helpStyle("Press any key to quit")
}
