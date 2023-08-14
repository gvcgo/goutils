package main

import (
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
	// f.SetUrl("https://golang.google.cn/dl/go1.21.0.linux-amd64.tar.gz")
	f.SetUrl("https://mirrors.aliyun.com/golang/go1.21.0.linux-amd64.tar.gz?spm=a2c6h.25603864.0.0.33337c45JOHx3F")
	// f.SetUrl("https://mirrors.nju.edu.cn/golang/go1.21.0.linux-amd64.tar.gz")
	// f.SetUrl("https://mirrors.ustc.edu.cn/golang/go1.21.0.linux-amd64.tar.gz")
	f.SetThreadNum(8)
	f.GetAndSaveFile(`C:\Users\moqsien\data\projects\go\src\goutils\go1.21.0.linux-amd64.tar.gz`, true)
}
