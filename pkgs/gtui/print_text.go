package gtui

import (
	"github.com/pterm/pterm"
)

func Green(v ...interface{}) {
	pterm.Println(pterm.Green(v...))
}

func Yellow(v ...interface{}) {
	pterm.Println(pterm.Yellow(v...))
}

func Cyan(v ...interface{}) {
	pterm.Println(pterm.Cyan(v...))
}

func Magenta(v ...interface{}) {
	pterm.Println(pterm.Magenta(v...))
}

func White(v ...interface{}) {
	pterm.Println(pterm.White(v...))
}

func Gray(v ...interface{}) {
	pterm.Println(pterm.Gray(v...))
}

func Blue(v ...interface{}) {
	pterm.Println(pterm.Blue(v...))
}
