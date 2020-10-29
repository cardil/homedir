package fs

import (
	"io/ioutil"

	"github.com/cardil/knative-homedir/internal/entity"
	"github.com/mitchellh/go-homedir"
)

// List will return a contents of a given directory.
func List(path string) (*entity.Listing, error) {
	fullpath, err := homedir.Expand(path)
	if err != nil {
		return nil, err
	}
	fileInfos, err := ioutil.ReadDir(fullpath)
	if err != nil {
		return nil, err
	}
	files := make([]entity.File, len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {
		fi := fileInfos[i]
		files[i] = entity.File{
			Name: fi.Name(),
			Mode: fi.Mode().String(),
			Size: int(fi.Size()),
		}
	}
	return &entity.Listing{
		Path:  fullpath,
		Files: files,
	}, nil
}
