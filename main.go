package main

import (
	"fmt"
	"os"

	"github.com/moqsien/goutils/pkgs/crypt"
)

func main() {
	if content, err := os.ReadFile("conf.txt"); err != nil {
		fmt.Println(err)
	} else {
		r, _ := crypt.DefaultCrypt.AesDecrypt(content)
		fmt.Println(string(r))
		// fmt.Println(err)
	}
}
