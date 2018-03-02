package storage

type LocalStorage struct {
	Path string
}

func (LocalStorage) Retrieve() {}
func (LocalStorage) Store()    {}
