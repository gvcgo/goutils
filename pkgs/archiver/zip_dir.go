package archiver

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	ezip "github.com/alexmullins/zip"
)

func (that *Archiver) SetZipName(zipName string) {
	that.ZipName = zipName
}

func (that *Archiver) GetZipPath() string {
	if that.ZipName == "" {
		name := filepath.Base(that.SrcFilePath)
		if strings.Contains(name, ".") {
			name = strings.Split(name, ".")[0]
		}
		that.ZipName = name + ".zip"
	}
	return filepath.Join(that.DstDir, that.ZipName)
}

func (that *Archiver) ZipDir() (err error) {
	if that.Password != "" {
		return that.zipWithPassword()
	}
	return that.zipOnly()
}

func (that *Archiver) zipOnly() (err error) {
	fz, err := os.Create(that.GetZipPath())
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

func (that *Archiver) zipWithPassword() (err error) {
	desc := that.GetZipPath()
	zipfile, err := os.Create(desc)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := ezip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(that.SrcFilePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := ezip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, filepath.Dir(that.SrcFilePath)+"/")
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Name = path[len(that.SrcFilePath)+1:]
			header.Method = zip.Deflate
		}

		header.SetPassword(that.Password)
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		var file *os.File
		if !info.IsDir() {
			file, err = os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
	return err
}

func ZipTest() {
	a, _ := NewArchiver(`/Users/moqsien/.ssh`, `/Volumes/Data/projects/go/src/goutils`)
	a.SetZipName("dotSSH.zip")
	a.SetPassword("123456")
	err := a.zipWithPassword()
	fmt.Println(err)
}
