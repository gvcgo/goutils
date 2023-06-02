package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/moqsien/goutils/pkgs/gtui"
	utils "github.com/moqsien/goutils/pkgs/gutils"
)

const (
	IsChildEnv     = "GOUTILS_IS_CHILD_PROCESS"
	IsChildProcess = "GOUTILS_IS_CHILD_PROCESS=true"
)

type Daemon struct {
	workdir string
	batName string
}

func NewDaemon() *Daemon {
	return &Daemon{}
}

// for windows
func (that *Daemon) SetWorkdir(d string) {
	that.workdir = d
	utils.MakeDirs(d)
}

// for windows
func (that *Daemon) SetScriptName(batName string) {
	if batName != "" && !strings.HasSuffix(batName, ".bat") {
		batName = fmt.Sprintf("%s.bat", batName)
	}
	that.batName = batName
}

func getWinScriptName() string {
	if fPath, err := os.Executable(); err == nil {
		name := strings.TrimSuffix(filepath.Base(fPath), ".exe")
		return fmt.Sprintf("%s.bat", name)
	}
	return "daemon_script.bat"
}

func (that *Daemon) getWinScriptPath(osArgs ...string) (fPath string) {
	if that.batName == "" {
		that.batName = getWinScriptName()
	}
	fPath = filepath.Join(that.workdir, that.batName)
	if ok, _ := utils.PathIsExist(fPath); !ok {
		batStr := strings.Join(osArgs, " ")
		os.WriteFile(fPath, []byte(batStr), os.ModePerm)
	}
	return fPath
}

func (that *Daemon) Run(osArgs ...string) {
	if len(osArgs) == 0 {
		panic("no osArgs specified")
	}
	if isChild := os.Getenv(IsChildEnv); isChild != "" {
		return
	}
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		batFilePath := that.getWinScriptPath(osArgs...)
		cmd = exec.Command("powershell", "Start-Process", "-WindowStyle", "hidden", "-FilePath", batFilePath)
	} else {
		cmd = exec.Command(osArgs[0], osArgs[1:]...)
	}
	cmd.Env = append(os.Environ(), IsChildProcess)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); err != nil {
		gtui.PrintErrorf("start %s failed, error: %v\n", osArgs[0], err)
		os.Exit(1)
	}
	gtui.PrintSuccessf("%s [PID] %d running...\n", osArgs[0], cmd.Process.Pid)
	os.Exit(0)
}
