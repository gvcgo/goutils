package remoter

type DavRemoter struct {
	Host     string
	UserName string
	Password string
}

func NewDavRemoter(host, username, password string) (d *DavRemoter) {
	d = &DavRemoter{
		Host:     host,
		UserName: username,
		Password: password,
	}
	return
}

func (that *DavRemoter) CreateRepo() (r string) {
	return
}

func (that *DavRemoter) UploadFile() (r string) {
	return
}

func (that *DavRemoter) DeleteFile() (r string) {
	return
}
