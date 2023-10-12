package archiver

import (
	"fmt"
	"os"
	"strings"

	archive "github.com/mholt/archiver/v3"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
	"github.com/moqsien/goutils/pkgs/gutils"
	"github.com/moqsien/xtractr"
)

type Archiver struct {
	DstDir      string
	SrcFilePath string
	ZipName     string
	UseAchiver  bool
	Password    string
}

func NewArchiver(srcFilePath string, dstDir string, useArchiver ...bool) (*Archiver, error) {
	if ok, _ := gutils.PathIsExist(srcFilePath); !ok {
		return nil, fmt.Errorf("srcfile path does not exists")
	}
	if ok, _ := gutils.PathIsExist(dstDir); !ok {
		os.MkdirAll(dstDir, os.ModePerm)
	}
	a := &Archiver{DstDir: dstDir, SrcFilePath: srcFilePath}
	if len(useArchiver) > 0 {
		a.UseAchiver = useArchiver[0]
	}
	return a, nil
}

func (that *Archiver) UnArchive() (string, error) {
	if strings.HasSuffix(that.SrcFilePath, ".tar.xz") {
		err := XZDecompress(that.SrcFilePath, that.DstDir)
		return that.DstDir, err
	}
	if that.UseAchiver {
		err := archive.Unarchive(that.SrcFilePath, that.DstDir)
		return that.DstDir, err
	}
	x := &xtractr.XFile{
		FilePath:  that.SrcFilePath,
		OutputDir: that.DstDir,
		FileMode:  os.ModePerm,
		DirMode:   os.ModePerm,
	}
	if that.Password != "" {
		x.Password = that.Password
	}
	// size is how many bytes were written.
	// files may be nil, but will contain any files written (even with an error).
	size, files, _, err := xtractr.ExtractFile(x)
	if files == nil || err != nil {
		gprint.PrintError("%v, %+v", size, err)
	}
	return that.DstDir, err
}

func (that *Archiver) SetPassword(p string) {
	that.Password = p
}

func ArchiverTest() {
	a, _ := NewArchiver(`C:\Users\moqsien\.gvc\typst_files\typst_x64_linux.tar.xz`, `C:\Users\moqsien\.gvc\typst_files\test`)
	_, err := a.UnArchive()
	fmt.Println(err)
}
