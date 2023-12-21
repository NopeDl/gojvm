package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absolutePath string
}

func newZipEntry(path string) *ZipEntry {
  absPath, err := filepath.Abs(path)
  if err != nil {
    panic(err)
  }
	return &ZipEntry{absolutePath: absPath}
}

// 重写接口方法
func (zipEntry *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
  rc, err := zip.OpenReader(zipEntry.absolutePath)
  if err != nil {
    return nil, nil, err
  }
  defer rc.Close()
  for _, file := range rc.File {
    if file.Name == className {
      r, fileError := file.Open()
      defer r.Close()
      if fileError != nil {
        return nil, nil, err
      }
      data, readError := io.ReadAll(r)
      return data, zipEntry, readError
    }
  }
  return nil, nil, errors.New("class not found: " + className)
}


// 重写接口方法
func (zipEntry *ZipEntry) ToString() string {
  return zipEntry.absolutePath
}
