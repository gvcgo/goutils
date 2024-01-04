package remoter

const (
	GithubAPI string = "https://api.github.com"
)

type GhRemoter struct {
	AuthToken string
}

func NewGhRemoter(authToken string) (g *GhRemoter) {
	g = &GhRemoter{
		AuthToken: authToken,
	}
	return
}

func (that *GhRemoter) CreateRepo() (r string) {
	return
}

func (that *GhRemoter) UploadFile() (r string) {
	return
}

func (that *GhRemoter) DeleteFile() (r string) {
	return
}
