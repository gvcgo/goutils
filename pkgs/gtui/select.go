package gtui

import "github.com/pterm/pterm"

type HandleSelected func(string) string

type Select struct {
	OptionList []string
	Handle     HandleSelected
}

func NewSelect(opts []string, handle HandleSelected) *Select {
	return &Select{OptionList: opts, Handle: handle}
}

func (that *Select) Start() string {
	selectedOpt, err := pterm.DefaultInteractiveSelect.WithOptions(that.OptionList).Show()
	if err != nil {
		PrintError(err)
		return ""
	}
	if that.Handle != nil {
		return that.Handle(selectedOpt)
	}
	return selectedOpt
}
