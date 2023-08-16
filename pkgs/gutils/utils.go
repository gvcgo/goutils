package gutils

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"

	"github.com/moqsien/goutils/pkgs/gtui"
)

func PathIsExist(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

func MakeDirs(dirs ...string) {
	for _, d := range dirs {
		if ok, _ := PathIsExist(d); !ok {
			if err := os.MkdirAll(d, os.ModePerm); err != nil {
				fmt.Println("mkdir failed: ", err)
			}
		}
	}
}

func Closeq(v any) {
	if c, ok := v.(io.Closer); ok {
		silently(c.Close())
	}
}

func silently(_ ...any) {}

func CopyFile(src, dst string) (written int64, err error) {
	srcFile, err := os.Open(src)

	if err != nil {
		gtui.PrintError(fmt.Sprintf("Cannot open file: %+v", err))
		return
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		gtui.PrintError(fmt.Sprintf("Cannot open file: %+v", err))
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

func CheckSum(fpath, cType, cSum string) (r bool) {
	if cSum != ComputeSum(fpath, cType) {
		gtui.PrintError("Checksum failed.")
		return
	}
	gtui.PrintSuccess("Checksum succeeded.")
	return true
}

func ComputeSum(fpath, sumType string) (sumStr string) {
	f, err := os.Open(fpath)
	if err != nil {
		gtui.PrintError(fmt.Sprintf("Open file failed: %+v", err))
		return
	}
	defer f.Close()

	var h hash.Hash
	switch strings.ToLower(sumType) {
	case "sha256":
		h = sha256.New()
	case "sha1":
		h = sha1.New()
	case "sha512":
		h = sha512.New()
	default:
		gtui.PrintError(fmt.Sprintf("[Crypto] %s is not supported.", sumType))
		return
	}

	if _, err = io.Copy(h, f); err != nil {
		gtui.PrintError(fmt.Sprintf("Copy file failed: %+v", err))
		return
	}

	sumStr = hex.EncodeToString(h.Sum(nil))
	return
}
