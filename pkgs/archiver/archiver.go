package archiver

import (
	"fmt"
	"os"
	"strings"

	"github.com/moqsien/goutils/pkgs/gtui"
	"github.com/moqsien/goutils/pkgs/gutils"
	"golift.io/xtractr"
)

type Archiver struct {
	DstDir      string
	SrcFilePath string
}

func NewArchiver(srcFilePath string, dstDir string) (*Archiver, error) {
	if ok, _ := gutils.PathIsExist(srcFilePath); !ok {
		return nil, fmt.Errorf("srcfile path does not exists")
	}
	if ok, _ := gutils.PathIsExist(dstDir); !ok {
		os.MkdirAll(dstDir, 0777)
	}
	return &Archiver{DstDir: dstDir, SrcFilePath: srcFilePath}, nil
}

func (that *Archiver) UnArchive() (string, error) {
	if strings.HasSuffix(that.SrcFilePath, ".tar.xz") {
		err := XZDecompress(that.SrcFilePath, that.DstDir)
		return that.DstDir, err
	}
	x := &xtractr.XFile{
		FilePath:  that.SrcFilePath,
		OutputDir: that.DstDir,
	}

	// size is how many bytes were written.
	// files may be nil, but will contain any files written (even with an error).
	size, files, _, err := xtractr.ExtractFile(x)
	if files == nil || err != nil {
		gtui.PrintError(size, err)
	}
	return that.DstDir, err
}

func ArchiverTest() {
	a, _ := NewArchiver(`C:\Users\moqsien\.gvc\typst_files\typst_x64_linux.tar.xz`, `C:\Users\moqsien\.gvc\typst_files\test`)
	_, err := a.UnArchive()
	fmt.Println(err)
}
