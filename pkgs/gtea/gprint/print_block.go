package gprint

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type BOption func(bp *BlockPrinter)

func WithWidth(width int) BOption {
	return func(bp *BlockPrinter) {
		bp.Style.Width(width)
	}
}

func WithHeight(height int) BOption {
	return func(bp *BlockPrinter) {
		bp.Style.Height(height)
	}
}

func WithPadding(i ...int) BOption {
	return func(bp *BlockPrinter) {
		bp.Style.Padding(i...)
	}
}

func WithForeground(color string) BOption {
	return func(bp *BlockPrinter) {
		bp.Style.Foreground(lipgloss.Color(color))
	}
}

func WithBackground(lightColor, darkColor string) BOption {
	return func(bp *BlockPrinter) {
		bp.Style.Background(lipgloss.AdaptiveColor{Light: lightColor, Dark: darkColor})
	}
}

func WithBold(bold bool) BOption {
	return func(bp *BlockPrinter) {
		bp.Style.Bold(bold)
	}
}

func WithItalic(italic bool) BOption {
	return func(bp *BlockPrinter) {
		bp.Style.Italic(italic)
	}
}

func WithAlign(position lipgloss.Position) BOption {
	return func(bp *BlockPrinter) {
		bp.Style.Align(position)
	}
}

type BlockPrinter struct {
	content string
	Style   *lipgloss.Style
}

func NewBlockPrinter(content string, opts ...BOption) (bp *BlockPrinter) {
	s := lipgloss.NewStyle().Margin(1, 3, 0, 0)
	bp = &BlockPrinter{
		Style:   &s,
		content: content,
	}
	for _, opt := range opts {
		opt(bp)
	}
	return
}

func (that *BlockPrinter) String() string {
	return that.Style.Render(that.content)
}

func (that *BlockPrinter) Println() {
	fmt.Println(that.String())
}

func PrintlnByDefault(content string) {
	bp := NewBlockPrinter(
		content,
		WithAlign(lipgloss.Left),
		WithForeground("#FAFAFA"),
		WithBackground("#874BFD", "#7D56F4"),
		WithPadding(2, 6),
		WithHeight(6),
		WithWidth(53),
		WithBold(true),
		WithItalic(true),
	)
	bp.Println()
}
