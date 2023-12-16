package bar

import (
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

type ProgressMsg float64

type ExtraMsg string

type ErrorMsg struct{ err error }

func finalPause() tea.Cmd {
	return tea.Tick(time.Millisecond*750, func(_ time.Time) tea.Msg {
		return nil
	})
}

type DownloadModel struct {
	progress Model
	err      error
	sweep    func()
}

func (dm *DownloadModel) SetSweep(sweep func()) {
	dm.sweep = sweep
}

func (dm *DownloadModel) Init() tea.Cmd {
	return nil
}

func (dm *DownloadModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			if dm.sweep != nil {
				dm.sweep()
			}
			os.Exit(1)
		}
		return dm, nil
	case tea.WindowSizeMsg:
		dm.progress.Width = msg.Width - padding*2 - 4
		if dm.progress.Width > maxWidth {
			dm.progress.Width = maxWidth
		}
		return dm, nil

	case ErrorMsg:
		dm.err = msg.err
		return dm, tea.Quit

	case ProgressMsg:
		var cmds []tea.Cmd

		if msg >= 1.0 {
			cmds = append(cmds, tea.Sequence(finalPause(), tea.Quit))
		}
		cmds = append(cmds, dm.progress.SetPercent(float64(msg)))
		return dm, tea.Batch(cmds...)
	case ExtraMsg:
		dm.progress.SetExtra(string(msg))
		return dm, nil
	// FrameMsg is sent when the progress bar wants to animate itself
	case FrameMsg:
		progressModel, cmd := dm.progress.Update(msg)
		dm.progress = progressModel.(Model)
		return dm, cmd

	default:
		return dm, nil
	}
}

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Render

func (dm *DownloadModel) View() string {
	if dm.err != nil {
		return "Error downloading: " + dm.err.Error() + "\n"
	}

	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + dm.progress.View() + "\n" +
		pad + helpStyle(`Press "q" to quit.`) + "\n"
}
