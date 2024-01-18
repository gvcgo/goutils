package storage

type IStorage interface {
	CreateRepo(repoName string) []byte
	GetRepoInfo(repoName string) []byte
	GetContents(repoName, remotePath, fileName string) []byte
	UploadFile(repoName, remotePath, localPath, shaStr string) []byte
	DeleteFile(repoName, remotePath, fileName, shaStr string) []byte
}
