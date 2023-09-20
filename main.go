package main

import (
	"fmt"
	"log"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/client"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/moqsien/goutils/pkgs/ggit/gssh"
	"github.com/moqsien/goutils/pkgs/gutils"
)

type Comparable int

func (that Comparable) Less(other gutils.IComparable) bool {
	i := other.(Comparable)
	return that < i
}

func cloneRepo(_url, dir, publicKeyPath string) (*git.Repository, error) {
	log.Printf("cloning %s into %s", _url, dir)
	auth, keyErr := publicKey(publicKeyPath)
	if keyErr != nil {
		return nil, keyErr
	}

	client.InstallProtocol("ssh", gssh.DefaultClient)
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		Progress:     os.Stdout,
		URL:          _url,
		Auth:         auth,
		ProxyOptions: transport.ProxyOptions{URL: "http://localhost:2023"},
	})

	if err != nil {
		log.Printf("clone git repo error: %s", err)
		return nil, err
	}

	return r, nil
}

func publicKey(filePath string) (*ssh.PublicKeys, error) {
	var publicKey *ssh.PublicKeys
	sshKey, _ := os.ReadFile(filePath)
	publicKey, err := ssh.NewPublicKeys("git", []byte(sshKey), "")
	// fmt.Println(sshKey)
	if err != nil {
		return nil, err
	}
	return publicKey, err
}

func main() {
	// if content, err := os.ReadFile("conf.txt"); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	r, _ := crypt.DefaultCrypt.AesDecrypt(content)
	// 	fmt.Println(string(r))
	// 	// fmt.Println(err)
	// }

	// f := request.NewFetcher()
	// f.SetUrl("https://golang.google.cn/dl/go1.21.0.linux-amd64.tar.gz")
	// f.SetUrl("https://mirrors.aliyun.com/golang/go1.21.0.linux-amd64.tar.gz?spm=a2c6h.25603864.0.0.33337c45JOHx3F")
	// f.SetUrl("https://mirrors.nju.edu.cn/golang/go1.21.0.linux-amd64.tar.gz")
	// f.SetUrl("https://mirrors.ustc.edu.cn/golang/go1.21.0.linux-amd64.tar.gz")
	// f.SetThreadNum(8)
	// f.GetAndSaveFile(`C:\Users\moqsien\data\projects\go\src\goutils\go1.21.0.linux-amd64.tar.gz`, true)
	// archiver.ArchiverTest()
	// uuid := gutils.NewUUID()
	// fmt.Println(uuid.String())
	// s, err := base64.RawStdEncoding.DecodeString("Y2RuLmFwcHNmbHllci5jJSXvv71bJe+/vR9JSXvvv70l77+9")
	// fmt.Println(string(s), err)

	// str := "abcdfafafjkjalfjkdfnan94385=+!f"
	// r := crypt.EncodeBase64(str)
	// fmt.Println(r)
	// rd := crypt.DecodeBase64(r)
	// fmt.Println(rd)

	// iList := []Comparable{6, 8, 2, 4, 1, 5, 7, 3}
	// cList := []gutils.IComparable{}
	// for _, i := range iList {
	// 	cList = append(cList, i)
	// }
	// gutils.QuickSort(cList, 0, len(iList)-1)
	// fmt.Println(cList)

	// a, _ := archiver.NewArchiver(`C:\Users\moqsien\data\projects\go\src\goutils\test`, `C:\Users\moqsien\data\projects\go\src\goutils`)
	// a.SetZipName("test.zip")
	// err := a.ZipDir()
	// fmt.Println(err)
	keyPath := `C:\Users\moqsien\.ssh\id_rsa`
	_, err := cloneRepo("git@github.com:moqsien/goktrl.git", `C:\Users\moqsien\data\projects\go\src\play\test`, keyPath)
	fmt.Println(err)
}
