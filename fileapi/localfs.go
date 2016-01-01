package fileapi

import (
	"errors"
	"os"
	"strings"

	"types"
)

type LocalFsEntry struct {
	types.FsEntry
	name string
}

func (lfe *LocalFsEntry) Name() string {
	return lfe.name
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

// List
func (LocalFS) List(fpath string) ([]LocalFsEntry, error) {
	p, e := rfsMangle(fpath)
	if e != nil {
		return nil, e
	}
	osFile, e := os.Open(p)
	if e != nil {
		return nil, e
	}

	// ディレクトリを読んで FsEntry 一覧を取得する
	fis, e := osFile.Readdir(0)
	rs := make([]LocalFsEntry, len(fis))
	for i, fi := range fis {
		rs[i].name = fi.Name()
	}

	return rs, nil
}
