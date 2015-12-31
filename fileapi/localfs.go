package fileapi

import (
	"errors"
	"os"
	"strings"
)

// type rdir
// ----------------------------------------------------------------
//
type rdir struct {
	d *os.File
}

func (d rdir) Readdir(count int) ([]NamedAttr, error) {
	fis, e := d.d.Readdir(count)
	if e != nil {
		return nil, e
	}
	rs := make([]NamedAttr, len(fis))
	for i, fi := range fis {
		rs[i].Name = fi.Name()
		// rs[i].FillFrom(fi)
	}
	return rs, nil
}
func (d rdir) Close() error {
	return d.d.Close()
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

// OpenDir
func (LocalFS) OpenDir(fpath string) (Dir, error) {
	p, e := rfsMangle(fpath)
	if e != nil {
		return nil, e
	}
	f, e := os.Open(p)
	if e != nil {
		return nil, e
	}
	return rdir{f}, nil
}
