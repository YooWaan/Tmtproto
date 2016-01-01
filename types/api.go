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
	FsEntry
	Readdir(count int) (map[string]string, error)
}

type File interface {
	FsEntry
	FStat() (map[string]string, error)
	FSetStat(map[string]string) error
}

type FileSystem interface {
	List(name string, flags uint32, attr map[string]string) (FsEntry, error)
	Remove(FsEntry) error
	Rename(old string, new string, flags uint32) error
	Mkdir(name string, attr map[string]string) error
	Read(FsEntry) (io.Reader, error)
	Write(FsEntry) (io.Writer, error)

	// TODO: Stat 関連を実装する
	// Stat(name string, islstat bool) (map[string]string, error)
	// SetStat(name string, attr map[string]string) error
}
