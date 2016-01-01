package fileapi

import (
	"errors"
	"path"

	"types"
)

var Failure = errors.New("Failure")

type EmptyFile struct {}

func (EmptyFile) Close() error                       { return nil }
func (EmptyFile) ReadAt([]byte, int64) (int, error)  { return 0, Failure }
func (EmptyFile) WriteAt([]byte, int64) (int, error) { return 0, Failure }
func (EmptyFile) FStat() (map[string]string, error)              { return nil, Failure }
func (EmptyFile) FSetStat(map[string]string) error               { return Failure }

type EmptyFS struct {}

func (EmptyFS) Open(string, uint32, map[string]string) (types.DirectoryEntry, error)  { return nil, Failure }
func (EmptyFS) Remove(string) error                           { return Failure }
func (EmptyFS) Rename(string, string, uint32) error           { return Failure }
func (EmptyFS) Mkdir(string, map[string]string) error                     { return Failure }
func (EmptyFS) Rmdir(string) error                            { return Failure }
func (EmptyFS) Stat(string, bool) (map[string]string, error)              { return nil, Failure }
func (EmptyFS) SetStat(string, map[string]string) error                   { return Failure }
func (EmptyFS) ReadLink(p string) (string, error)             { return "", Failure }
func (EmptyFS) CreateLink(p string, t string, f uint32) error { return Failure }
func (EmptyFS) RealPath(p string) (string, error)             { return simpleRealPath(p), nil }

func simpleRealPath(fp string) string {
	switch fp {
	case "", ".":
		fp = "/"
	default:
		fp = path.Clean(fp)
	}
	return fp
}
