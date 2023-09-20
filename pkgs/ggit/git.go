package ggit

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/client"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/moqsien/goutils/pkgs/ggit/gssh"
	"github.com/moqsien/goutils/pkgs/gtui"
)

type Git struct {
	ProxyUrl   string
	SSHKeyPath string
}

func NewGit(proxyUrl string) (g *Git) {
	return &Git{
		ProxyUrl: proxyUrl,
	}
}

func (that *Git) SetSSHKeyPath(keyPath string) {
	that.SSHKeyPath = keyPath
}

func (that *Git) parseProjectNameFromUrl(projectUrl string) (name string) {
	sList := strings.Split(projectUrl, "/")
	if len(sList) == 0 {
		return
	}
	name = sList[len(sList)-1]
	name = strings.TrimRight(name, ".git")
	return
}

func (that *Git) getSSHKey() (*ssh.PublicKeys, error) {
	if that.SSHKeyPath == "" {
		homeDir, _ := os.UserHomeDir()
		that.SSHKeyPath = filepath.Join(homeDir, ".ssh", "id_rsa")
	}

	if that.SSHKeyPath == "" {
		return nil, fmt.Errorf("can not find ssh key in: %s", that.SSHKeyPath)
	}
	var publicKey *ssh.PublicKeys
	sshKey, _ := os.ReadFile(that.SSHKeyPath)
	publicKey, err := ssh.NewPublicKeys("git", []byte(sshKey), "")
	if err != nil {
		return nil, err
	}
	return publicKey, err
}

func (that *Git) installPortocol(projectUrl string) {
	client.InstallProtocol("ssh", gssh.DefaultClient)
}

func (that *Git) Clone(projectUrl string) (*git.Repository, error) {
	projectName := that.parseProjectNameFromUrl(projectUrl)
	if projectName == "" {
		gtui.PrintError("can not fine project name!")
		return nil, fmt.Errorf("can not fine project name: %s", projectUrl)
	}
	cwdir, err := os.Getwd()
	if err != nil {
		gtui.PrintError(err)
		return nil, err
	}
	auth, err := that.getSSHKey()
	if err != nil {
		return nil, err
	}
	that.installPortocol(projectUrl)
	r, err := git.PlainClone(filepath.Join(cwdir, projectName), false, &git.CloneOptions{
		Progress:     os.Stdout,
		URL:          projectUrl,
		Auth:         auth,
		ProxyOptions: transport.ProxyOptions{URL: "http://localhost:2023"},
	})

	if err != nil {
		gtui.PrintErrorf("clone git repo error: %s", err)
		return nil, err
	}
	return r, nil
}
