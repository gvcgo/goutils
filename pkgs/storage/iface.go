package storage

type IStorage interface {
	CreateRepo(string) []byte
	GetRepoInfo(string) []byte
	GetContents(string, string, string) []byte
	UploadFile(string, string, string, string) []byte
	DeleteFile(string, string, string, string) []byte
}
