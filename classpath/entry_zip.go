package classpath

import (
	"archive/zip"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir: absDir}
}
func (d *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(d.absDir)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, d, nil
		}

	}

}

func (d *ZipEntry) String() string {
	return d.absDir
}
