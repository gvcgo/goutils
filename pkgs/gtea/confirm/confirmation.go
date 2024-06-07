package confirm

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/gvcgo/goutils/pkgs/gtea/gprint"
)

type ConfirmOption func(cfm *Confirmation)

func WithPrompt(prompt string, style ...lipgloss.Style) ConfirmOption {

	return func(cfm *Confirmation) {
		if len(style) == 0 {
			cfm.confirm.Prompt = gprint.CyanStr(prompt)
		} else {
			cfm.confirm.Prompt = style[0].Render(prompt)
		}
	}
}

type Confirmation struct {
	confirm confirmation.Confirmation
	ok      bool
}

func NewConfirmation(options ...ConfirmOption) *Confirmation {
	cfm := &Confirmation{
		confirm: *confirmation.New("Yes or No?", confirmation.NewValue(false)),
	}

	cfm.confirm.Template = confirmation.TemplateYN
	cfm.confirm.ResultTemplate = confirmation.ResultTemplateYN
	for _, option := range options {
		option(cfm)
	}

	return cfm
}

func (cfm *Confirmation) Run() {
	cfm.ok, _ = cfm.confirm.RunPrompt()
}

func (cfm *Confirmation) Result() bool {
	return cfm.ok
}
