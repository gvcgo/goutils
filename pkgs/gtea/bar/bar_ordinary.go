package bar

import (
	"fmt"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gvcgo/goutils/pkgs/gtea/program"
)

type OrdinaryBar struct {
	Program         *tea.Program
	model           *DownloadModel
	total           int64
	processed       int64
	succeeded       int64
	enableSucceeded bool
	lock            *sync.Mutex
}

func NewOrdinaryBar(opts ...Option) (bar *OrdinaryBar) {
	model := &DownloadModel{
		progress: New(opts...),
	}
	bar = &OrdinaryBar{
		Program: tea.NewProgram(model),
		model:   model,
		lock:    &sync.Mutex{},
	}
	return
}

func (bar *OrdinaryBar) SetProgramOpts(opts ...tea.ProgramOption) {
	if bar.model != nil {
		bar.Program = tea.NewProgram(bar.model, opts...)
	}
}

func (bar *OrdinaryBar) SetTotal(total int64) {
	bar.total = total
}

func (bar *OrdinaryBar) EnableSucceeded() {
	bar.enableSucceeded = true
}

func (bar *OrdinaryBar) prepareExtraInfo() (extra string) {
	if bar.enableSucceeded {
		extra = fmt.Sprintf("%d/%d|succeeded: %d", bar.processed, bar.total, bar.succeeded)
	} else {
		extra = fmt.Sprintf("%d/%d", bar.processed, bar.total)
	}
	return
}

func (bar *OrdinaryBar) Add(processed, succeeded int) {
	bar.lock.Lock()
	bar.processed += int64(processed)
	bar.succeeded += int64(succeeded)
	if bar.total > 0 && bar.Program != nil {
		ratio := float64(bar.processed) / float64(bar.total)
		bar.Program.Send(ProgressMsg(ratio))
		bar.Program.Send(ExtraMsg(bar.prepareExtraInfo()))
	}
	bar.lock.Unlock()
}

func (bar *OrdinaryBar) AddOnlyProcessed(processed int) {
	bar.lock.Lock()
	bar.processed += int64(processed)
	if bar.total > 0 && bar.Program != nil {
		ratio := float64(bar.processed) / float64(bar.total)
		bar.Program.Send(ProgressMsg(ratio))
		bar.Program.Send(ExtraMsg(bar.prepareExtraInfo()))
	}
	bar.lock.Unlock()
}

func (bar *OrdinaryBar) Run() {
	program.Run(bar.Program)
}
