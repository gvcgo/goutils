package gutils

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	mrand "math/rand"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gvcgo/goutils/pkgs/gtea/gprint"
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
				gprint.PrintError("mkdir failed: %+v", err)
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

func CopyFile(src, dst string) (written int64, err error) {
	srcFile, err := os.Open(src)

	if err != nil {
		gprint.PrintError(fmt.Sprintf("Cannot open file: %+v", err))
		return
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		gprint.PrintError(fmt.Sprintf("Cannot open file: %+v", err))
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

// CopyAFile copies the file at the source path to the provided destination.
func CopyAFile(source, destination string) error {
	//Validate the source and destination paths
	if len(source) == 0 {
		return errors.New("you must provide a source file path")
	}

	if len(destination) == 0 {
		return errors.New("you must provide a destination file path")
	}

	//Verify the source path refers to a regular file
	sourceFileInfo, err := os.Lstat(source)
	if err != nil {
		return err
	}

	//Handle regular files differently than symbolic links and other non-regular files.
	if sourceFileInfo.Mode().IsRegular() {
		//open the source file
		sourceFile, err := os.Open(source)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		//create the destinatin file
		destinationFile, err := os.Create(destination)
		if err != nil {
			return err
		}
		defer destinationFile.Close()

		//copy the source file contents to the destination file
		if _, err = io.Copy(destinationFile, sourceFile); err != nil {
			return err
		}

		//replicate the source file mode for the destination file
		if err := os.Chmod(destination, sourceFileInfo.Mode()); err != nil {
			return err
		}
	} else if sourceFileInfo.Mode()&os.ModeSymlink != 0 {
		linkDestinaton, err := os.Readlink(source)
		if err != nil {
			return errors.New("Unable to read symlink. " + err.Error())
		}

		if err := os.Symlink(linkDestinaton, destination); err != nil {
			return errors.New("Unable to replicate symlink. " + err.Error())
		}
	} else {
		return errors.New("Unable to use io.Copy on file with mode " + gconv.String(sourceFileInfo.Mode()))
	}

	return nil
}

// CopyDirectory copies the directory at the source path to the provided destination, with the option of recursively copying subdirectories.
func CopyDirectory(source string, destination string, recursive bool) error {
	if len(source) == 0 || len(destination) == 0 {
		return errors.New("file paths must not be empty")
	}

	//get properties of the source directory
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	//create the destination directory
	err = os.MkdirAll(destination, sourceInfo.Mode())
	if err != nil {
		return err
	}

	sourceDirectory, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceDirectory.Close()

	objects, err := sourceDirectory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, object := range objects {
		if object.Name() == ".Trashes" || object.Name() == ".DS_Store" {
			continue
		}

		sourceObjectName := source + string(filepath.Separator) + object.Name()
		destObjectName := destination + string(filepath.Separator) + object.Name()

		if object.IsDir() {
			//create sub-directories
			err = CopyDirectory(sourceObjectName, destObjectName, true)
			if err != nil {
				return err
			}
		} else {
			err = CopyAFile(sourceObjectName, destObjectName)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CheckSum(fpath, cType, cSum string) (r bool) {
	if cSum != ComputeSum(fpath, cType) {
		gprint.PrintError("Checksum failed.")
		return
	}
	gprint.PrintSuccess("Checksum succeeded.")
	return true
}

func ComputeSum(fpath, sumType string) (sumStr string) {
	f, err := os.Open(fpath)
	if err != nil {
		gprint.PrintError(fmt.Sprintf("Open file failed: %+v", err))
		return
	}
	defer f.Close()

	var h hash.Hash
	switch strings.ToLower(sumType) {
	case "sha256":
		h = sha256.New()
	case "sha1":
		h = sha1.New()
	case "sha512":
		h = sha512.New()
	default:
		gprint.PrintError(fmt.Sprintf("[Crypto] %s is not supported.", sumType))
		return
	}

	if _, err = io.Copy(h, f); err != nil {
		gprint.PrintError(fmt.Sprintf("Copy file failed: %+v", err))
		return
	}

	sumStr = hex.EncodeToString(h.Sum(nil))
	return
}

func VerifyUrls(rawUrl string) (r bool) {
	r = true
	_, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		r = false
		return
	}
	url, err := url.Parse(rawUrl)
	if err != nil || url.Scheme == "" || url.Host == "" {
		r = false
		return
	}
	return
}

const (
	Win     string = "win"
	Zsh     string = "zsh"
	Bash    string = "bash"
	Windows string = "windows"
	Darwin  string = "darwin"
	Linux   string = "linux"
)

func GetHomeDir() string {
	home, _ := os.UserHomeDir()
	return home
}

func GetShell() (shell string) {
	if runtime.GOOS == Windows {
		return Win
	}
	s := os.Getenv("SHELL")
	if strings.Contains(s, "zsh") {
		return Zsh
	}
	return Bash
}

func GetShellRcFile() (rc string) {
	shell := GetShell()
	switch shell {
	case Zsh:
		rc = filepath.Join(GetHomeDir(), ".zshrc")
	case Bash:
		rc = filepath.Join(GetHomeDir(), ".bashrc")
	default:
		rc = Win
	}
	return
}

func FlushPathEnvForUnix() (err error) {
	if runtime.GOOS != Windows {
		err = exec.Command("source", GetShellRcFile()).Run()
	}
	return
}

func ExecuteSysCommand(collectOutput bool, workDir string, args ...string) (*bytes.Buffer, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == Windows {
		args = append([]string{"/c"}, args...)
		cmd = exec.Command("cmd", args...)
	} else {
		FlushPathEnvForUnix()
		cmd = exec.Command(args[0], args[1:]...)
	}
	cmd.Env = os.Environ()
	var output bytes.Buffer
	if collectOutput {
		cmd.Stdout = &output
	} else {
		cmd.Stdout = os.Stdout
	}
	if workDir != "" {
		cmd.Dir = workDir
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return &output, cmd.Run()
}

type UUID [16]byte

// create a new uuid v4
func NewUUID() *UUID {
	u := &UUID{}
	_, err := rand.Read(u[:16])
	if err != nil {
		panic(err)
	}

	u[8] = (u[8] | 0x80) & 0xBf
	u[6] = (u[6] | 0x40) & 0x4f
	return u
}

func (u *UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func RandomString(slength int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, slength)
	for i := range s {
		s[i] = letters[mrand.Intn(len(letters))]
	}
	return string(s)
}
