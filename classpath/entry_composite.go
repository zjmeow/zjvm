package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var entries []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entries = append(entries, newEntry(path))
	}
	return entries
}
func (c CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, e, err := entry.ReadClass(className)
		if err == nil {
			return data, e, err
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (c CompositeEntry) String() string {

}
