package gprint

import (
	"fmt"
)

func Green(format string, v ...any) {
	fmt.Println(GreenStr(format, v...))
}

func Yellow(format string, v ...any) {
	fmt.Println(YellowStr(format, v...))
}

func Cyan(format string, v ...any) {
	fmt.Println(CyanStr(format, v...))
}

func Magenta(format string, v ...any) {
	fmt.Println(MagentaStr(format, v...))
}

func White(format string, v ...any) {
	fmt.Println(WhiteStr(format, v...))
}

func Gray(format string, v ...any) {
	fmt.Println(GrayStr(format, v...))
}

func Blue(format string, v ...any) {
	fmt.Println(BlueStr(format, v...))
}

func Pink(format string, v ...any) {
	fmt.Println(PinkStr(format, v...))
}

func Brown(format string, v ...any) {
	fmt.Println(BrownStr(format, v...))
}

func Rose(format string, v ...any) {
	fmt.Println(RoseStr(format, v...))
}

func Red(format string, v ...any) {
	fmt.Println(RedStr(format, v...))
}

func Orange(format string, v ...any) {
	fmt.Println(OrangeStr(format, v...))
}
