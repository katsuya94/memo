package storage

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/katsuya94/memo/util"
)

type LocalStorage struct {
	Path string
}

func (s LocalStorage) Get(d util.Date) (util.Memo, error) {
	f, err := os.Open(path.Join(s.Path, d.String()))
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	memo, err := util.ReadMemo(f)
	if err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	return memo, nil
}

func (s LocalStorage) Put(d util.Date, memo util.Memo) error {
	if err := os.MkdirAll(s.Path, 0700); err != nil {
		return err
	}

	f, err := os.OpenFile(path.Join(s.Path, d.String()), os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	if err := util.WriteMemo(memo, f); err != nil {
		return err
	}

	return f.Close()
}

func (s LocalStorage) List() ([]util.Date, error) {
	var dates []util.Date
	fileinfos, err := ioutil.ReadDir(s.Path)
	if err != nil {
		return dates, err
	}
	for _, fileinfo := range fileinfos {
		dates = append(dates, util.NewDateFromString(fileinfo.Name()))
	}
	return dates, nil
}
