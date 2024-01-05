package main

import (
	"github.com/moqsien/goutils/pkgs/gutils"
	"github.com/moqsien/goutils/pkgs/storage"
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
	// g.CloneBySSH("git@github.com:jesseduffield/lazygit.git")

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

	// itemList := selector.NewItemList()
	// itemList.Add("win", "windows")
	// itemList.Add("linux", "linux")
	// itemList.Add("mac", "darwin")
	// // sel := selector.NewSelector(itemList, selector.WithShowStatusBar(true), selector.WithTitle("Choose OS type:"), selector.WithEnbleInfinite(true), selector.WidthEnableMulti(true))
	// sel := selector.NewSelector(itemList, selector.WithShowStatusBar(true), selector.WithTitle("Choose OS type:"), selector.WithEnbleInfinite(true))
	// sel.Run()
	// fmt.Println(sel.Value())

	// s := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
	// 	"Name: gvc",
	// 	"Version: v1.5.6(4e189a)",
	// 	"UpdateAt: Thu Sep 21 12:53:09 2023 +0800",
	// 	"Homepage: https://github.com/moqsien/gvc",
	// 	"Email: moqsien2022@gmail.com",
	// )
	// gprint.PrintlnByDefault(s)

	// columns := []gtable.Column{
	// 	{Title: "Rank", Width: 10},
	// 	{Title: "City", Width: 20},
	// 	{Title: "Country", Width: 20},
	// 	{Title: "Population", Width: 50},
	// }

	// rows := []gtable.Row{
	// 	{"1", "Tokyo", "Japan", "37,274,000"},
	// 	{"2", "Delhi", "India", "32,065,760"},
	// 	{"3", "Shanghai", "China", "28,516,904"},
	// 	{"4", "Dhaka", "Bangladesh", "22,478,116"},
	// 	{"5", "São Paulo", "Brazil", "22,429,800"},
	// 	{"6", "Mexico City", "Mexico", "22,085,140"},
	// }

	// t := gtable.NewTable(gtable.WithColumns(columns), gtable.WithRows(rows), gtable.WithHeight(7), gtable.WithFocused(true))
	// t.Run()

	// style := lipgloss.NewStyle().Width(20).MaxWidth(20).Inline(true)
	// renderedCell := style.Render(runewidth.Truncate("hello test", 20, "*"))
	// hStyle := lipgloss.NewStyle().Bold(true).Padding(0, 1).BorderStyle(lipgloss.NormalBorder()).
	// 	BorderForeground(lipgloss.Color("240")).
	// 	BorderBottom(true).
	// 	Bold(false)
	// // fmt.Println(hStyle.Render(renderedCell))
	// var s = make([]string, 0, len(columns))
	// for _, col := range columns {
	// 	style := lipgloss.NewStyle().Width(col.Width).MaxWidth(col.Width).Inline(true)
	// 	renderedCell := style.Render(runewidth.Truncate(col.Title, col.Width, "…"))
	// 	s = append(s, hStyle.Render(renderedCell))
	// 	// fmt.Println(hStyle.Render(renderedCell))
	// }
	// fmt.Println(gtable.JoinHorizontal(lipgloss.Left, s...))
	// cfm := confirm.NewConfirm(confirm.WithTitle("Do you want to have something?"))
	// cfm.Run()
	// fmt.Println(cfm.Result())

	// mInput := input.NewMultiInput()
	// mInput.AddOneItem("url", input.MWithWidth(60))
	// mInput.AddOneItem("username", input.MWithWidth(60))
	// mInput.AddOneItem("password", input.MWithWidth(60), input.MWithEchoMode(textinput.EchoPassword), input.MWithEchoChar("*"))
	// mInput.AddOneOption("gpt_model", []string{"hello", "golang", "test"}, input.MWithWidth(60))
	// mInput.Run()
	// fmt.Printf("%+v\n", mInput.Values())

	// w := &sync.WaitGroup{}
	// for i := 0; i < 1000; i++ {
	// 	w.Add(1)
	// 	var v int = i
	// 	go func() {
	// 		gprint.PrintError("test: %d", v)
	// 		w.Done()
	// 	}()
	// }
	// w.Wait()

	// 	git := ggit.NewGit()
	// 	git.SetWorkDir(`C:\Users\moqsien\data\projects\go\src\goutils`)
	// 	git.ShowLatestTag()
	// archiver.ArchiverTest()
	// archiver.ZipTest()

	// values := []string{
	// 	"test1",
	// 	"hello2",
	// 	"golang3",
	// }
	// opt := input.NewOption(values, input.WithCharlimit(100))
	// opt.Run()
	// fmt.Println(opt.Value())

	// storage.GhTest()
	storage.GtTest()
}
