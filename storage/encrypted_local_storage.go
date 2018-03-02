package storage

type EncryptedLocalStorage struct {
	Path string
}

func (EncryptedLocalStorage) Retrieve() {}
func (EncryptedLocalStorage) Store()    {}
