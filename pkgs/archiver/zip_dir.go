package archiver

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func (that *Archiver) SetZipName(zipName string) {
	that.ZipName = zipName
}

func (that *Archiver) ZipDir() (err error) {
	if that.ZipName == "" {
		name := filepath.Base(that.SrcFilePath)
		if strings.Contains(name, ".") {
			name = strings.Split(name, ".")[0]
		}
		that.ZipName = name + ".zip"
	}
	fz, err := os.Create(filepath.Join(that.DstDir, that.ZipName))
	if err != nil {
		return err
	}
	defer fz.Close()

	w := zip.NewWriter(fz)
	defer w.Close()

	err = filepath.Walk(that.SrcFilePath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fDest, err := w.Create(path[len(that.SrcFilePath)+1:])
			if err != nil {
				return err
			}
			fSrc, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fSrc.Close()
			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}
