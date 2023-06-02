package gtui

import (
	"github.com/pterm/pterm"
)

func PrintSuccess(v ...any) {
	pterm.Success.Println(v...)
}

func PrintSuccessf(format string, v ...any) {
	pterm.Success.Printfln(format, v...)
}

func PrintError(v ...any) {
	pterm.Error.Println(v...)
}

func PrintErrorf(format string, v ...any) {
	pterm.Error.Printfln(format, v...)
}

func PrintInfo(v ...any) {
	pterm.Info.Println(v...)
}

func PrintInfof(format string, v ...any) {
	pterm.Info.Printfln(format, v...)
}

func PrintWarning(v ...any) {
	pterm.Warning.Println(v...)
}

func PrintWarningf(format string, v ...any) {
	pterm.Warning.Printfln(format, v...)
}

func PrintFatal(v ...any) {
	pterm.Fatal.WithFatal(false).Println(v...)
}

func PrintFatalf(format string, v ...any) {
	pterm.Fatal.WithFatal(false).Printfln(format, v...)
}
