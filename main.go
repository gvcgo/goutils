package main

import (
	"fmt"

	"github.com/moqsien/goutils/pkgs/gtea/selector"
	"github.com/moqsien/goutils/pkgs/gutils"
)

type Comparable int

func (that Comparable) Less(other gutils.IComparable) bool {
	i := other.(Comparable)
	return that < i
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
	// g := ggit.NewGit()
	// g.SetProxyUrl("http://localhost:2023")
	// g.CloneBySSH("git@github.com:moqsien/goktrl.git")
	// g.AddTagAndPushToRemote("v1.3.9")
	// g.DeleteTagAndPushToRemote("v1.3.9")
	// err := g.CommitAndPush("update")
	// fmt.Println(err)
	// gtea.Run("https://gitlab.com/moqsien/gvc_resources/-/raw/main/gvc_windows-amd64.zip")
	// gtea.TestDownload("https://gitlab.com/moqsien/gvc_resources/-/raw/main/gvc_windows-amd64.zip")

	// f := request.NewFetcher()
	// f.SetUrl("https://gitlab.com/moqsien/gvc_resources/-/raw/main/gvc_windows-amd64.zip")
	// f.SetThreadNum(2)
	// f.GetAndSaveFile("gvc_windows-amd64.zip", true)

	// gtui.PrintWarning("hello")
	// gtui.PrintInfo("hello")
	// gtui.PrintError("hello")
	// gtui.PrintFatal("hello")
	// gtui.PrintSuccess("hello")

	// gprint.PrintWarning("hello")
	// gprint.PrintInfo("hello")
	// gprint.PrintError("hello")
	// gprint.PrintFatal("hello")
	// gprint.PrintSuccess("hello")

	// gprint.Green("hello")
	// gprint.Yellow("hello")
	// gprint.Cyan("hello")
	// gprint.Magenta("hello")
	// gprint.White("hello")
	// gprint.Gray("hello")
	// gprint.Blue("hello")
	// gprint.Pink("hello")
	// gprint.Brown("hello")
	// gprint.Rose("hello")
	// gprint.Red("hello")
	// gprint.Orange("hello")

	// content := "hello hello hello hello hello hello hello hello hello hello hello hello hello hello hello "
	// s := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4")).Render(content)
	// fmt.Println(s)

	// fcolor := gprint.NewFadeColors(content)
	// fcolor.Println()

	// ipt := input.NewInput(input.WithEchoMode(textinput.EchoPassword), input.WithEchoChar("*"), input.WithPlaceholder("password"))
	// ipt.Run()
	// fmt.Println(ipt.Value())

	itemList := selector.NewItemList()
	itemList.Add("win", "windows")
	itemList.Add("linux", "linux")
	itemList.Add("mac", "darwin")
	sel := selector.NewSelector(itemList, selector.WithShowStatusBar(true), selector.WithTitle("Choose OS type:"), selector.WithEnbleInfinite(true))
	sel.Run()
	fmt.Println(sel.Value())
}
