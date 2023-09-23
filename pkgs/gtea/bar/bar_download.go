package bar

import (
	"fmt"
	"io"
	"os"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
)

type DownloadBar struct {
	Program    *tea.Program
	total      int64
	downloaded int64
	lock       *sync.Mutex
}

func NewDownloadBar(opts ...Option) (bar *DownloadBar) {
	model := &DownloadModel{
		progress: New(opts...),
	}
	bar = &DownloadBar{
		Program: tea.NewProgram(model),
		lock:    &sync.Mutex{},
	}
	return
}

func (bar *DownloadBar) SetTotal(total int64) {
	bar.total = total
}

func (bar *DownloadBar) prepareExtraInfo() (extra string) {
	mbSize := 1048576
	kbSize := 1024
	if bar.total > int64(mbSize) {
		extra = fmt.Sprintf(
			" [%.2f/%.2f MB]",
			float64(bar.downloaded)/float64(mbSize),
			float64(bar.total)/float64(mbSize),
		)
	} else {
		extra = fmt.Sprintf(
			" [%.2f/%.2f KB]",
			float64(bar.downloaded)/float64(kbSize),
			float64(bar.total)/float64(kbSize),
		)
	}
	return
}

func (bar *DownloadBar) Write(p []byte) (int, error) {
	bar.lock.Lock()
	bar.downloaded += int64(len(p))
	if bar.total > 0 && bar.Program != nil {
		ratio := float64(bar.downloaded) / float64(bar.total)
		bar.Program.Send(ProgressMsg(ratio))
		bar.Program.Send(ExtraMsg(bar.prepareExtraInfo()))
	}
	bar.lock.Unlock()
	return len(p), nil
}

func (bar *DownloadBar) Copy(bodyReader io.Reader, storageFile *os.File) {
	// TeeReader calls bar.Write() each time a new packet is received
	_, err := io.Copy(storageFile, io.TeeReader(bodyReader, bar))
	if err != nil {
		bar.Program.Send(ErrorMsg{err})
	}
}

func (bar *DownloadBar) Run() error {
	if _, err := bar.Program.Run(); err != nil {
		return err
	}
	return nil
}
