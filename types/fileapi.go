package types

import (
	"io"
	// "os"
)

type FsEntry interface {
	Name() string
	IsDir() bool
	IsFile() bool
}

type Dir interface {
	FsEntry
	Readdir(count int) (map[string]string, error)
}

type File interface {
	FsEntry
}

type FileSystem interface {
	List(name string) (FsEntry, error)
	Remove(FsEntry) error
	Rename(old string, new string) error
	Mkdir(name string) error
	Read(FsEntry) (io.Reader, error)
	Write(FsEntry) (io.Writer, error)

	// TODO: Stat 関連を実装する
	// Stat(name string, islstat bool) (map[string]string, error)
	// SetStat(name string, attr map[string]string) error
}
