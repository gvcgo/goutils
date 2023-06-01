package gutils

import (
	"fmt"
	"io"
	"os"
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
