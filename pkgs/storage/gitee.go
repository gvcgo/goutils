package storage

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/moqsien/goutils/pkgs/gtea/gprint"
	"github.com/moqsien/goutils/pkgs/gutils"
	"github.com/moqsien/goutils/pkgs/request"
)

const (
	GiteeAPI string = "https://gitee.com/api/v5"
)

/*
Docs: https://gitee.com/api/v5/swagger#/postV5UserRepos
*/
type GtStorage struct {
	UserName  string
	AuthToken string
	fetcher   *request.Fetcher
}

func NewGtStorage(username, authToken string) (g *GtStorage) {
	g = &GtStorage{
		UserName:  username,
		AuthToken: authToken,
		fetcher:   request.NewFetcher(),
	}
	return
}

func (that *GtStorage) CreateRepo(repoName string) (r []byte) {
	// https://gitee.com/api/v5/user/repos
	that.fetcher.SetUrl(fmt.Sprintf("%s/user/repos", GiteeAPI))
	that.fetcher.PostBody = map[string]interface{}{
		"access_token": that.AuthToken,
		"name":         repoName,
		"auto_init":    true,
	}
	if resp := that.fetcher.Post(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	// set repo to be public.
	r = that.PatchRepo(repoName)
	return
}

func (that *GtStorage) PatchRepo(repoName string) (r []byte) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s", GiteeAPI, that.UserName, repoName))
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s", GiteeAPI, that.UserName, repoName))
	that.fetcher.PostBody = map[string]interface{}{
		"access_token": that.AuthToken,
		"name":         repoName,
		"owner":        that.UserName,
		"repo":         "",
		"private":      false,
	}
	if resp := that.fetcher.Patch(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func (that *GtStorage) GetRepoInfo(repoName string) (r []byte) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}
	that.fetcher.SetUrl(fmt.Sprintf(
		"%s/repos/%s/%s?access_token=%s",
		GiteeAPI,
		that.UserName,
		repoName,
		that.AuthToken,
	))
	that.fetcher.Timeout = 60 * time.Second
	if resp := that.fetcher.Get(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func (that *GtStorage) GetContents(repoName, remotePath, fileName string) (r []byte) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}/contents(/{path})?access_token=xxx
	remotePath = strings.TrimLeft(filepath.Join(remotePath, fileName), "/")
	that.fetcher.SetUrl(fmt.Sprintf(
		"%s/repos/%s/%s/contents/%s?access_token=%s",
		GiteeAPI,
		that.UserName,
		repoName,
		remotePath,
		that.AuthToken,
	))
	that.fetcher.Timeout = 60 * time.Second
	if resp := that.fetcher.Get(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func (that *GtStorage) UploadFile(repoName, remotePath, localPath, shaStr string) (r []byte) {
	if ok, _ := gutils.PathIsExist(localPath); !ok {
		gprint.PrintError("file: %s does not exist.", localPath)
		return
	}
	fName := filepath.Base(localPath)
	result := that.GetContents(repoName, remotePath, fName)

	flagStr := `"download_url":`
	if !strings.Contains(string(result), flagStr) {
		r = that.CreateFile(repoName, remotePath, localPath)
	} else {
		r = that.UpdateFile(repoName, remotePath, localPath, shaStr)
	}
	return
}

func (that *GtStorage) CreateFile(repoName, remotePath, localPath string) (r []byte) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}/contents/{path}
	fName := filepath.Base(localPath)
	remotePath = strings.TrimLeft(filepath.Join(remotePath, fName), "/")
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s/contents/%s", GiteeAPI, that.UserName, repoName, remotePath))
	that.fetcher.Timeout = 30 * time.Minute

	content, _ := os.ReadFile(localPath)
	that.fetcher.PostBody = map[string]interface{}{
		"access_token": that.AuthToken,
		"message":      fmt.Sprintf("update file: %s.", fName),
		"content":      base64.StdEncoding.EncodeToString(content),
	}
	if resp := that.fetcher.Post(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func (that *GtStorage) UpdateFile(repoName, remotePath, localPath, shaStr string) (r []byte) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}/contents/{path}
	fName := filepath.Base(localPath)
	remotePath = strings.TrimLeft(filepath.Join(remotePath, fName), "/")
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s/contents/%s", GiteeAPI, that.UserName, repoName, remotePath))
	that.fetcher.Timeout = 30 * time.Minute

	content, _ := os.ReadFile(localPath)
	that.fetcher.PostBody = map[string]interface{}{
		"access_token": that.AuthToken,
		"owner":        that.UserName,
		"repo":         repoName,
		"path":         remotePath,
		"message":      fmt.Sprintf("update file: %s.", fName),
		"content":      base64.StdEncoding.EncodeToString(content),
		"sha":          shaStr,
	}
	if resp := that.fetcher.Put(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func (that *GtStorage) DeleteFile(repoName, remotePath, fileName, shaStr string) (r []byte) {
	// https://gitee.com/api/v5/repos/{owner}/{repo}/contents/{path}
	remotePath = strings.TrimLeft(filepath.Join(remotePath, fileName), "/")
	that.fetcher.SetUrl(fmt.Sprintf("%s/repos/%s/%s/contents/%s", GiteeAPI, that.UserName, repoName, remotePath))
	that.fetcher.Timeout = 30 * time.Minute
	that.fetcher.PostBody = map[string]interface{}{
		"access_token": that.AuthToken,
		"owner":        that.UserName,
		"repo":         repoName,
		"path":         remotePath,
		"message":      fmt.Sprintf("delete file: %s.", fileName),
		"sha":          shaStr,
	}
	if resp := that.fetcher.Delete(); resp != nil {
		defer resp.RawResponse.Body.Close()
		r, _ = io.ReadAll(resp.RawResponse.Body)
	}
	return
}

func GtTest() {
	key := "xxx"
	user := "moqsien"
	repoName := "gvc_conf_test"
	gtr := NewGtStorage(user, key)
	gtr.CreateRepo(repoName)
	// r := gtr.PatchRepo(repoName)
	// r := gtr.GetRepoInfo(repoName)
	r := gtr.GetContents(repoName, "", "test2.txt")
	// j := gjson.New(r)
	// shaStr := j.Get("sha").String()
	// localPath := "/Volumes/data/projects/go/src/goutils/test2.txt"
	// r = gtr.UploadFile(repoName, "", localPath, shaStr)

	// r = gtr.DeleteFile(repoName, "", "LICENSE", shaStr)
	fmt.Println(string(r))
	// j := gjson.New(r)
	// fmt.Println(j.Get("download_url").String())
}
