package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/moqsien/goutils/pkgs/gtea/gtable"
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

	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}

	rows := []table.Row{
		{"1", "Tokyo", "Japan", "37,274,000"},
		{"2", "Delhi", "India", "32,065,760"},
		{"3", "Shanghai", "China", "28,516,904"},
		{"4", "Dhaka", "Bangladesh", "22,478,116"},
		{"5", "São Paulo", "Brazil", "22,429,800"},
		{"6", "Mexico City", "Mexico", "22,085,140"},
		{"7", "Cairo", "Egypt", "21,750,020"},
		{"8", "Beijing", "China", "21,333,332"},
		{"9", "Mumbai", "India", "20,961,472"},
		{"10", "Osaka", "Japan", "19,059,856"},
		{"11", "Chongqing", "China", "16,874,740"},
		{"12", "Karachi", "Pakistan", "16,839,950"},
		{"13", "Istanbul", "Turkey", "15,636,243"},
		{"14", "Kinshasa", "DR Congo", "15,628,085"},
		{"15", "Lagos", "Nigeria", "15,387,639"},
		{"16", "Buenos Aires", "Argentina", "15,369,919"},
		{"17", "Kolkata", "India", "15,133,888"},
		{"18", "Manila", "Philippines", "14,406,059"},
		{"19", "Tianjin", "China", "14,011,828"},
		{"20", "Guangzhou", "China", "13,964,637"},
		{"21", "Rio De Janeiro", "Brazil", "13,634,274"},
		{"22", "Lahore", "Pakistan", "13,541,764"},
		{"23", "Bangalore", "India", "13,193,035"},
		{"24", "Shenzhen", "China", "12,831,330"},
		{"25", "Moscow", "Russia", "12,640,818"},
		{"26", "Chennai", "India", "11,503,293"},
		{"27", "Bogota", "Colombia", "11,344,312"},
		{"28", "Paris", "France", "11,142,303"},
		{"29", "Jakarta", "Indonesia", "11,074,811"},
		{"30", "Lima", "Peru", "11,044,607"},
		{"31", "Bangkok", "Thailand", "10,899,698"},
	}

	t := gtable.NewTable(table.WithColumns(columns), table.WithRows(rows), table.WithHeight(7), table.WithFocused(true))
	t.Run()
}
