package remoter

const (
	GithubAPI           string = "https://api.github.com"
	AcceptHeader        string = "application/vnd.github.v3+json"
	AuthorizationHeader string = "token %s"
)

type GhRemoter struct {
	UserName  string
	AuthToken string
	Proxy     string
}

func NewGhRemoter(username, authToken string) (g *GhRemoter) {
	g = &GhRemoter{
		UserName:  username,
		AuthToken: authToken,
	}
	return
}

func (that *GhRemoter) CreateRepo(repoName string) (r string) {
	return
}

func (that *GhRemoter) UploadFile() (r string) {
	return
}

func (that *GhRemoter) GetFileInfo() (r string) {
	return
}

func (that *GhRemoter) DeleteFile() (r string) {
	return
}
