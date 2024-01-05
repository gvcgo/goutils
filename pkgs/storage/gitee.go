package storage

import "github.com/moqsien/goutils/pkgs/request"

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

func (that *GtStorage) CreateRepo() (r []byte) {
	return
}

func (that *GtStorage) GetRepoInfo() (r []byte) {
	return
}

func (that *GtStorage) GetContents() (r []byte) {
	return
}

func (that *GtStorage) UploadFile() (r []byte) {
	return
}

func (that *GtStorage) DeleteFile() (r []byte) {
	return
}
