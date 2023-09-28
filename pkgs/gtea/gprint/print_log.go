package gprint

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	hintColor = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000"))
	txtColor  = lipgloss.NewStyle()
)

const (
	HintInfo    string = "  INFO   "
	HintSuccess string = " SUCCESS "
	HintWarning string = " WARNING "
	HintError   string = "  ERROR  "
	HintFatal   string = "  FATAL  "
)

func prepareHint(hint, backgroudColor string) string {
	return hintColor.Copy().Background(lipgloss.Color(backgroudColor)).Render(hint)
}

func PrintInfo(format string, v ...any) {
	fmt.Println(prepareHint(HintInfo, "#40E0D0") + " " + txtColor.Copy().Foreground(lipgloss.Color("#40E0D0")).Render(fmt.Sprintf(format, v...)))
}

func PrintSuccess(format string, v ...any) {
	fmt.Println(prepareHint(HintSuccess, "#32CD32") + " " + txtColor.Copy().Foreground(lipgloss.Color("#32CD32")).Render(fmt.Sprintf(format, v...)))
}

func PrintWarning(format string, v ...any) {
	fmt.Println(prepareHint(HintWarning, "#FFFF00") + " " + txtColor.Copy().Foreground(lipgloss.Color("#FFFF00")).Render(fmt.Sprintf(format, v...)))
}

func PrintError(format string, v ...any) {
	fmt.Println(prepareHint(HintError, "#FF6347") + " " + txtColor.Copy().Foreground(lipgloss.Color("#FF6347")).Render(fmt.Sprintf(format, v...)))
}

func PrintFatal(format string, v ...any) {
	fmt.Println(prepareHint(HintFatal, "#FF0000") + " " + txtColor.Copy().Foreground(lipgloss.Color("#FF0000")).Render(fmt.Sprintf(format, v...)))
}
