package classpath

import (
	"os"
	"strings"
)

const osPathSeperator = string(os.PathSeparator)
const osPathListSeperator = string(os.PathListSeparator)

type Entry interface {
	ReadClass(className string) ([]byte, Entry, error)
	ToString() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, osPathListSeperator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
