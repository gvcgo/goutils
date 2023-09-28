package gprint

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func GreenStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#32CD32")).Render(fmt.Sprintf(format, v...))
}

func YellowStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#FFFF00")).Render(fmt.Sprintf(format, v...))
}

func CyanStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#00FFFF")).Render(fmt.Sprintf(format, v...))
}

func MagentaStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#FF00FF")).Render(fmt.Sprintf(format, v...))
}

func WhiteStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#FFFFFF")).Render(fmt.Sprintf(format, v...))
}

func GrayStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#808080")).Render(fmt.Sprintf(format, v...))
}

func BlueStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#00BFFF")).Render(fmt.Sprintf(format, v...))
}

func PinkStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#FF69B4")).Render(fmt.Sprintf(format, v...))
}

func BrownStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#D2691E")).Render(fmt.Sprintf(format, v...))
}

func RoseStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#FFE4E1")).Render(fmt.Sprintf(format, v...))
}

func RedStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#FF0000")).Render(fmt.Sprintf(format, v...))
}

func OrangeStr(format string, v ...any) string {
	return txtColor.Copy().Foreground(lipgloss.Color("#FFA500")).Render(fmt.Sprintf(format, v...))
}
