package remoter

type GtRemoter struct {
	AuthToken string
}

func NewGtRemoter(authToken string) (g *GtRemoter) {
	g = &GtRemoter{
		AuthToken: authToken,
	}
	return
}

func (that *GtRemoter) CreateRepo() (r string) {
	return
}

func (that *GtRemoter) UploadFile() (r string) {
	return
}

func (that *GtRemoter) DeleteFile() (r string) {
	return
}
