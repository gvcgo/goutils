package storage

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/moqsien/goutils/pkgs/gtea/gprint"
	"github.com/moqsien/goutils/pkgs/gutils"
	"github.com/moqsien/goutils/pkgs/request"
)

const (
	GithubAPI           string = "https://api.github.com"
	AcceptHeader        string = "application/vnd.github.v3+json"
	AuthorizationHeader string = "token %s"
)

type GhStorage struct {
	UserName  string
	AuthToken string
	Proxy     string
	fetcher   *request.Fetcher
}

func NewGhStorage(username, authToken string) (g *GhStorage) {
	g = &GhStorage{
		UserName:  username,
		AuthToken: authToken,
		fetcher:   request.NewFetcher(),
	}
	g.initiate()
	return
}

func (that *GhStorage) initiate() {
	if that.Proxy != "" {
		that.fetcher.Proxy = that.Proxy
	}
	that.fetcher.Headers = map[string]string{
		"Accept":        AcceptHeader,
		"Authorization": fmt.Sprintf(AuthorizationHeader, that.AuthToken),
	}
}

func (that *GhStorage) CreateRepo(repoName string) (r []byte) {
	// https://api.github.com/user/repos
	that.fetcher.SetUrl(fmt.Sprintf("%s/%s", GithubAPI, "user/repos"))
	that.fetcher.PostBody = map[string]interface{}{
		"name": repoName,
	}
	that.fetcher.Timeout = 60 * time.Second
	if resp := that.fetcher.Post(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func (that *GhStorage) GetRepoInfo(repoName string) (r []byte) {
	// https://api.github.com/repos/{user}/{repo}
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s", GithubAPI, that.UserName, repoName))
	that.fetcher.Timeout = 60 * time.Second
	if resp := that.fetcher.Get(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func (that *GhStorage) GetContents(repoName, remotePath, fileName string) (r []byte) {
	// https://api.github.com/repos/{user}/{repo}/contents/
	remotePath = strings.TrimLeft(filepath.Join(remotePath, fileName), "/")
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s/contents/%s", GithubAPI, that.UserName, repoName, remotePath))
	that.fetcher.Timeout = 60 * time.Second
	if resp := that.fetcher.Get(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

/*
SHA is needed for Update or Delete not Create.
*/
func (that *GhStorage) UploadFile(repoName, remotePath, localPath, shaStr string) (r []byte) {
	if ok, _ := gutils.PathIsExist(localPath); !ok {
		gprint.PrintError("file: %s does not exist.", localPath)
		return
	}
	// https://api.github.com/repos/{user}/{repo}/contents/{path}/{filename}
	fName := filepath.Base(localPath)
	remotePath = strings.TrimLeft(filepath.Join(remotePath, fName), "/")
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s/contents/%s", GithubAPI, that.UserName, repoName, remotePath))
	that.fetcher.Timeout = 30 * time.Minute

	content, _ := os.ReadFile(localPath)
	that.fetcher.PostBody = map[string]interface{}{
		"message": fmt.Sprintf("update file: %s.", fName),
		"content": base64.StdEncoding.EncodeToString(content),
		"sha":     shaStr,
	}
	if resp := that.fetcher.Put(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func (that *GhStorage) GetFileInfo(repoName, remotePath, fileName string) (r []byte) {
	// https://api.github.com/repos/{user}/{repo}/contents/{path}/{filename}
	remotePath = strings.TrimLeft(filepath.Join(remotePath, fileName), "/")
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s/contents/%s", GithubAPI, that.UserName, repoName, remotePath))
	that.fetcher.Timeout = 30 * time.Minute
	if resp := that.fetcher.Get(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func (that *GhStorage) DeleteFile(repoName, remotePath, fileName, shaStr string) (r []byte) {
	// https://api.github.com/repos/{user}/{repo}/contents/{path}/{filename}
	remotePath = strings.TrimLeft(filepath.Join(remotePath, fileName), "/")
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s/contents/%s", GithubAPI, that.UserName, repoName, remotePath))
	that.fetcher.Timeout = 30 * time.Minute
	that.fetcher.PostBody = map[string]interface{}{
		"message": fmt.Sprintf("delete file: %s.", fileName),
		"sha":     shaStr,
	}
	if resp := that.fetcher.Delete(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func GhTest() {
	key := "xxx"
	user := "moqsien"
	proxyURI := "http://127.0.0.1:2023"
	repoName := "gvc_conf_test"
	ghr := NewGhStorage(user, key)
	ghr.Proxy = proxyURI

	// localPath := "/Users/moqsien/data/projects/go/src/gvc/build/g_darwin-amd64.zip"
	// r := ghr.CreateRepo(repoName)
	// r := ghr.GetRepoInfo(repoName)
	// r := ghr.UploadFile(repoName, "", localPath, "")
	// r := ghr.GetFileInfo(repoName, "", localPath)
	// r := ghr.GetContents(repoName, "", "")
	r := ghr.GetFileInfo(repoName, "", "g_darwin-amd64.zip")
	j := gjson.New(r)
	shaStr := j.Get("sha").String()
	r = ghr.DeleteFile(repoName, "", "g_darwin-amd64.zip", shaStr)
	fmt.Println(string(r))
}
