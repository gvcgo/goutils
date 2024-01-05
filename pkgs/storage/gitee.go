package storage

type GtStorage struct {
	AuthToken string
}

func NewGtStorage(authToken string) (g *GtStorage) {
	g = &GtStorage{
		AuthToken: authToken,
	}
	return
}

func (that *GtStorage) CreateRepo() (r string) {
	return
}

func (that *GtStorage) UploadFile() (r string) {
	return
}

func (that *GtStorage) DeleteFile() (r string) {
	return
}
