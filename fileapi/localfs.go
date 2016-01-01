package fileapi

import (
	"errors"
	"os"
	"strings"

	"types"
)

// type rdir
// ----------------------------------------------------------------
//
type rdir struct {
	d *os.File
}

func (rd rdir) Readdir(count int) (map[int]types.DirectoryEntry, error) {
	// TODO: ディレクトリ内を読み込む動作を記述する
	// fInfos, e := rd.d.Readdir(count)
	// if e != nil {
	// 	return nil, e
	// }
	// // rs := map[int]ret{}
	// rs := make(map[int]types.DirectoryEntry, len(fInfos))
	// for i, fInfo := range fInfos {
	// 	// rs[i] = ret{Name: fInfo.Name()}
	// }
	// return rs, nil
	return nil, nil
}
func (rd rdir) Close() error {
	return rd.d.Close()
}

// type LocalFS
// ----------------------------------------------------------------
//
type LocalFS struct {
	EmptyFS
}

// Warning:
// Use your own path mangling functionality in production code.
// This can be quite non-trivial depending on the operating system.
// The code below is not sufficient for production servers.
func rfsMangle(fpath string) (string, error) {
	if strings.Contains(fpath, "..") {
		return "<invalid>", errors.New("Invalid path")
	}
	if len(fpath) > 0 && fpath[0] == '/' {
		fpath = fpath[1:]
	}
	// fpath = "/tmp/test-sftpd/" + fpath
	return fpath, nil
}

// Open
func (LocalFS) Open(fpath string) (types.DirectoryEntry, error) {
	// TODO: ディレクトリ内を読み込む動作を記述する
	// p, e := rfsMangle(fpath)
	// if e != nil {
	// 	return nil, e
	// }
	// osFile, e := os.Open(p)
	// if e != nil {
	// 	return nil, e
	// }
	// return rdir{osFile}, nil
	return nil, nil
}
