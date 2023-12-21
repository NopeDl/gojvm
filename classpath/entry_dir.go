package classpath

import (
	"os"
	"path/filepath"
)

type DirEntry struct {
	absolutePath string
}

func newDirEntry(path string) *DirEntry {
  absPath, err := filepath.Abs(path)
  if err != nil {
    panic(err)
  }
	return &DirEntry{absolutePath: absPath}
}

// 重写接口方法
func (dirEntry *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
  joinPath := filepath.Join(dirEntry.absolutePath, className)
  byteArray, err := os.ReadFile(joinPath)
	return byteArray, dirEntry, err
}

// 重写接口方法
func (dirEntry *DirEntry) ToString() string {
  return dirEntry.absolutePath
}
