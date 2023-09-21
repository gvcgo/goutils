package ggit

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/client"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/moqsien/goutils/pkgs/ggit/ghttp"
	"github.com/moqsien/goutils/pkgs/ggit/gssh"
	"github.com/moqsien/goutils/pkgs/gtui"
	"github.com/moqsien/goutils/pkgs/gutils"
)

type Git struct {
	ProxyUrl   string
	SSHKeyPath string
}

func NewGit() (g *Git) {
	return &Git{}
}

func (that *Git) SetSSHKeyPath(keyPath string) {
	that.SSHKeyPath = keyPath
}

func (that *Git) SetProxyUrl(proxyUrl string) {
	that.ProxyUrl = proxyUrl
}

func (that *Git) getUsernameAndEmail() (userName string, email string) {
	buf, err := gutils.ExecuteSysCommand(true, ".", "git", "config", "user.name")
	if err == nil {
		content, _ := io.ReadAll(buf)
		userName = string(content)
	}
	buf, err = gutils.ExecuteSysCommand(true, ".", "git", "config", "user.email")
	if err == nil {
		content, _ := io.ReadAll(buf)
		email = string(content)
	}
	return strings.TrimSpace(userName), strings.TrimSpace(email)
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

func (that *Git) CloneBySSH(projectUrl string) (*git.Repository, error) {
	if !strings.HasPrefix(projectUrl, "git@") {
		gtui.PrintErrorf("unsupported scheme: %s", projectUrl)
		return nil, fmt.Errorf("unsupported scheme: %s", projectUrl)
	}
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
	client.InstallProtocol("ssh", gssh.DefaultClient)
	client.InstallProtocol("https", ghttp.DefaultClient)

	r, err := git.PlainClone(filepath.Join(cwdir, projectName), false, &git.CloneOptions{
		Progress:     os.Stdout,
		URL:          projectUrl,
		Auth:         auth,
		ProxyOptions: transport.ProxyOptions{URL: that.ProxyUrl},
	})
	if err != nil {
		gtui.PrintErrorf("clone git repo error: %s", err)
		return nil, err
	}
	return r, nil
}

func (that *Git) PullBySSH() error {
	cwdir, err := os.Getwd()
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	r, err := git.PlainOpen(cwdir)
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	auth, err := that.getSSHKey()
	if err != nil {
		return err
	}
	client.InstallProtocol("ssh", gssh.DefaultClient)
	client.InstallProtocol("https", ghttp.DefaultClient)

	err = w.Pull(&git.PullOptions{
		Progress:     os.Stdout,
		RemoteName:   "origin",
		Auth:         auth,
		ProxyOptions: transport.ProxyOptions{URL: that.ProxyUrl},
	})
	if err != nil {
		gtui.PrintError(err)
	}
	return err
}

func (that *Git) push(r *git.Repository, auth transport.AuthMethod, tag string) error {
	client.InstallProtocol("ssh", gssh.DefaultClient)
	client.InstallProtocol("https", ghttp.DefaultClient)
	var tagRef string
	switch tag {
	case "*":
		tagRef = "refs/tags/*:refs/tags/*"
	case "":
		tagRef = ""
	default:
		tagRef = fmt.Sprintf(":refs/tags/%s", tag)
	}
	po := &git.PushOptions{
		RemoteName:   "origin",
		Progress:     os.Stdout,
		Auth:         auth,
		ProxyOptions: transport.ProxyOptions{URL: that.ProxyUrl},
	}
	if tagRef != "" {
		po.RefSpecs = []config.RefSpec{config.RefSpec(tagRef)}
		po.FollowTags = true
	}
	err := r.Push(po)
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			gtui.PrintWarning("origin remote was up to date, no push done")
			return nil
		}
		gtui.PrintErrorf("push to remote origin error: %s", err)
		return err
	}
	return nil
}

func (that *Git) PushBySSH() error {
	cwdir, err := os.Getwd()
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	r, err := git.PlainOpen(cwdir)
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	auth, err := that.getSSHKey()
	if err != nil {
		return err
	}
	return that.push(r, auth, "")
}

func (that *Git) handleRenameError(w *git.Worktree, fPath string, err error) {
	errStr := err.Error()
	if strings.Contains(errStr, "rename") && strings.HasSuffix(errStr, "Access is denied.") {
		sList := strings.Split(errStr, " ")
		if len(sList) > 3 {
			os.RemoveAll(strings.Trim(sList[2], ":"))
		}
		w.Add(fPath)
	}
}

func (that *Git) handleNewFiles(w *git.Worktree, cwdir string) {
	status, err := w.Status()
	if err != nil {
		gtui.PrintError(err)
		return
	}
	sList := strings.Split(status.String(), "\n")
	for _, pStr := range sList {
		fmt.Println(pStr)
		if strings.HasPrefix(pStr, "?? ") {
			p := strings.TrimPrefix(pStr, "?? ")
			_, err := w.Add(p)
			that.handleRenameError(w, p, err)
		}
	}
}

func (that *Git) CommitAndPush(commitMsg string) error {
	cwdir, err := os.Getwd()
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	r, err := git.PlainOpen(cwdir)
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	auth, err := that.getSSHKey()
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	that.handleNewFiles(w, cwdir)
	name, email := that.getUsernameAndEmail()

	commit, err := w.Commit(commitMsg, &git.CommitOptions{
		All: true,
		Author: &object.Signature{
			Name:  name,
			Email: email,
			When:  time.Now(),
		},
	})

	if err != nil {
		return err
	}

	obj, err := r.CommitObject(commit)
	if err != nil {
		gtui.PrintError(err)
		return err
	}
	fmt.Println(obj)
	return that.push(r, auth, "")
}

func (that *Git) setTag(r *git.Repository, tag string) (bool, error) {
	gtui.PrintInfof("Set tag %s", tag)
	h, err := r.Head()
	if err != nil {
		gtui.PrintErrorf("get HEAD error: %s", err)
		return false, err
	}
	name, email := that.getUsernameAndEmail()
	_, err = r.CreateTag(tag, h.Hash(), &git.CreateTagOptions{
		Message: tag,
		Tagger: &object.Signature{
			When:  time.Now(),
			Name:  name,
			Email: email,
		},
	})
	if err != nil {
		gtui.PrintErrorf("create tag error: %s", err)
		return false, err
	}
	return true, nil
}

func (that *Git) AddTagAndPushToRemote(tag string) error {
	cwdir, err := os.Getwd()
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	r, err := git.PlainOpen(cwdir)
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	auth, err := that.getSSHKey()
	if err != nil {
		return err
	}
	_, err = that.setTag(r, tag)
	if err != nil {
		gtui.PrintErrorf("create tag error: %s", err)
		return err
	}
	return that.push(r, auth, "*")
}

func (that *Git) DeleteTagAndPushToRemote(tag string) error {
	gtui.PrintInfof("Delete tag %s", tag)
	cwdir, err := os.Getwd()
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	r, err := git.PlainOpen(cwdir)
	if err != nil {
		gtui.PrintError(err)
		return err
	}

	auth, err := that.getSSHKey()
	if err != nil {
		return err
	}

	err = r.DeleteTag(tag)
	if err != nil {
		gtui.PrintErrorf("delete local tag failed: %+v", err)
	}
	return that.push(r, auth, tag)
}
