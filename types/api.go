package types

import (
	"io"
	// "os"
)

type DirectoryEntry interface {
	io.Closer
	Name() string
	isDir() bool
	isFile() bool
}

type Dir interface {
	io.Closer
	Readdir(count int) (map[string]string, error)
}

type File interface {
	io.Closer
	io.ReaderAt
	io.WriterAt
	FStat() (map[string]string, error)
	FSetStat(map[string]string) error
}

type FileSystem interface {
	Open(name string, flags uint32, attr map[string]string) (DirectoryEntry, error)
	Remove(name string) error
	Rename(old string, new string, flags uint32) error
	Mkdir(name string, attr map[string]string) error
	Rmdir(name string) error
	Stat(name string, islstat bool) (map[string]string, error)
	SetStat(name string, attr map[string]string) error
}
