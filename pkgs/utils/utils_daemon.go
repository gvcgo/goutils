package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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
	MakeDirs(d)
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

func (that *Daemon) getWinScriptPath() (fPath string) {
	if that.batName == "" {
		that.batName = getWinScriptName()
	}
	fPath = filepath.Join(that.workdir, that.batName)
	if ok, _ := PathIsExist(fPath); !ok {
		batStr := strings.Join(os.Args, " ")
		os.WriteFile(fPath, []byte(batStr), os.ModePerm)
	}
	return fPath
}

func (that *Daemon) Run() {
	if isChild := os.Getenv(IsChildEnv); isChild != "" {
		return
	}
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		batFilePath := that.getWinScriptPath()
		cmd = exec.Command("powershell", "Start-Process", "-WindowStyle", "hidden", "-FilePath", batFilePath)
	} else {
		cmd = exec.Command(os.Args[0], os.Args[1:]...)
	}
	cmd.Env = append(os.Environ(), IsChildProcess)
	if err := cmd.Start(); err != nil {
		fmt.Printf("start %s failed, error: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	fmt.Printf("%s [PID] %d running...\n", os.Args[0], cmd.Process.Pid)
	os.Exit(0)
}
