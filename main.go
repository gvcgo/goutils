package main

import (
	"runtime"

	"github.com/moqsien/goutils/pkgs/request"
)

func main() {
	// if content, err := os.ReadFile("conf.txt"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	r, _ := crypt.DefaultCrypt.AesDecrypt(content)
	// 	fmt.Println(string(r))
	// 	// fmt.Println(err)
	// }

	f := request.NewFetcher()
	f.SetUrl("https://mirrors.ustc.edu.cn/golang/go1.21.0.linux-amd64.tar.gz")
	f.SetThreadNum(runtime.NumCPU())
	f.Download(`C:\Users\moqsien\data\projects\go\src\goutils\go1.21.0.linux-amd64.tar.gz`, true)
}
