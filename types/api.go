package types

import (
	"io"
	// "os"
)

type FsEntry interface {
	Name() string
	isDir() bool
	isFile() bool
}

type Dir interface {
	Readdir(count int) (map[string]string, error)
}

type File interface {
	FStat() (map[string]string, error)
	FSetStat(map[string]string) error
}

type FileSystem interface {
	List(name string, flags uint32, attr map[string]string) (FsEntry, error)
	Remove(name string) error
	Rename(old string, new string, flags uint32) error
	Mkdir(name string, attr map[string]string) error
	Read(FsEntry) Reader
	Write(FsEntry) Writer, error

	// TODO: Stat 関連を実装する
	// Stat(name string, islstat bool) (map[string]string, error)
	// SetStat(name string, attr map[string]string) error
}
