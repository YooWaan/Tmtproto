package fileapi

import (
	"io"
	// "os"
)

type NamedAttr struct {
	Name string
	Attr map[string]string
}

const (
	ATTR_SIZE    = 0x00000001
	ATTR_UIDGID  = 0x00000002
	ATTR_MODE    = 0x00000004
	ATTR_TIME    = 0x00000008
	MODE_REGULAR = 0100000
	MODE_DIR     = 0040000
)

type Dir interface {
	io.Closer
	Readdir(count int) ([]NamedAttr, error)
}

type File interface {
	io.Closer
	io.ReaderAt
	io.WriterAt
	FStat() (map[string]string, error)
	FSetStat(map[string]string) error
}

type FileSystem interface {
	OpenFile(name string, flags uint32, attr map[string]string) (File, error)
	OpenDir(name string) (Dir, error)
	Remove(name string) error
	Rename(old string, new string, flags uint32) error
	Mkdir(name string, attr map[string]string) error
	Rmdir(name string) error
	Stat(name string, islstat bool) (map[string]string, error)
	SetStat(name string, attr map[string]string) error
}
