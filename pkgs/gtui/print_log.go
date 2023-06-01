package gtui

import (
	"github.com/pterm/pterm"
)

func PrintSuccess(v ...any) {
	pterm.Success.Println(v...)
}

func SPrintSuccess(format string, v ...any) {
	pterm.Success.Sprintfln(format, v...)
}

func PrintError(v ...any) {
	pterm.Error.Println(v...)
}

func SPrintErrorf(format string, v ...any) {
	pterm.Error.Sprintfln(format, v...)
}

func PrintInfo(v ...any) {
	pterm.Error.Println(v...)
}

func SPrintInfof(format string, v ...any) {
	pterm.Info.Sprintfln(format, v...)
}

func PrintWarning(v ...any) {
	pterm.Warning.Println(v...)
}

func SPrintWarningf(format string, v ...any) {
	pterm.Warning.Sprintfln(format, v...)
}

func PrintFatal(v ...any) {
	pterm.Fatal.WithFatal(false).Println(v...)
}

func SPrintFatalf(format string, v ...any) {
	pterm.Fatal.WithFatal(false).Printfln(format, v...)
}
