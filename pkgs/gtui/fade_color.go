package gtui

import (
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pterm/pterm"
)

type FadeColors struct {
	content any
}

func NewFadeColors(content any) *FadeColors {
	return &FadeColors{
		content: content,
	}
}

func (that *FadeColors) Println() {
	from := pterm.NewRGB(0, 255, 255)
	to := pterm.NewRGB(255, 0, 255)
	var (
		fadeInfo string
		strs     []string
	)
	if res, ok := that.content.(string); ok {
		strs = strings.Split(res, "")
	} else if res, ok := that.content.([]string); ok {
		strs = strings.Split(strings.Join(res, "  "), "")
	} else {
		strs = strings.Split(gconv.String(that.content), "")
	}
	length := len(strs)
	for i := 0; i < length; i++ {
		fadeInfo += from.Fade(0, float32(length), float32(i), to).Sprint(strs[i])
	}
	pterm.Println(fadeInfo)
}
