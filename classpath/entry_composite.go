package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, osPathListSeperator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.ReadClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) ToString() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.ToString()
	}
	return strings.Join(strs, osPathListSeperator)
}
