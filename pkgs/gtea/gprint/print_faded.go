package gprint

import (
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/termenv"
)

type FadeColors struct {
	ColorA       *colorful.Color
	ColorB       *colorful.Color
	colorProfile termenv.Profile
	content      any
}

func NewFadeColors(content any) *FadeColors {
	return &FadeColors{
		content:      content,
		colorProfile: termenv.ColorProfile(),
	}
}

func (that *FadeColors) SetRange(colorA, colorB string) {
	a, _ := colorful.Hex(colorA)
	that.ColorA = &a
	b, _ := colorful.Hex(colorB)
	that.ColorB = &b
}

func (that *FadeColors) SetDefaultRange() {
	that.SetRange("#00FFFF", "#FF00FF")
}

func (that *FadeColors) String() string {
	if that.ColorA == nil || that.ColorB == nil {
		that.SetDefaultRange()
	}
	var (
		result *strings.Builder = &strings.Builder{}
		strs   []string
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
		c := (*that.ColorA).BlendLuv(*that.ColorB, float64(i)/float64(length)).Hex()
		result.WriteString(termenv.
			String(strs[i]).
			Foreground(that.colorProfile.Color(c)).
			String(),
		)
	}
	return result.String()
}

func (that *FadeColors) Println() {
	fmt.Println(that.String())
}
