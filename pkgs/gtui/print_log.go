package gtui

import (
	"github.com/pterm/pterm"
)

func PrintSuccess(v ...any) {
	pterm.Success.Println(v...)
}

func PrintSuccessf(format string, v ...any) {
	pterm.Success.Printf(format, v...)
}

func PrintError(v ...any) {
	pterm.Error.Println(v...)
}

func PrintErrorf(format string, v ...any) {
	pterm.Error.Printf(format, v...)
}

func PrintInfo(v ...any) {
	pterm.Info.Println(v...)
}

func PrintInfof(format string, v ...any) {
	pterm.Info.Printf(format, v...)
}

func PrintWarning(v ...any) {
	pterm.Warning.Println(v...)
}

func PrintWarningf(format string, v ...any) {
	pterm.Warning.Printf(format, v...)
}

func PrintFatal(v ...any) {
	pterm.Fatal.WithFatal(false).Println(v...)
}

func PrintFatalf(format string, v ...any) {
	pterm.Fatal.WithFatal(false).Printf(format, v...)
}
