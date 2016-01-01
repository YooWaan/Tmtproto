package fileapi

import (
	"os"
	// "errors"
	// "strings"

	"types"
)

type LocalFileSystemEntry struct {
	types.FsEntry
	name string
	osFileInfo os.FileInfo
}

func (lfe *LocalFileSystemEntry) Name() string {
	return lfe.name
}

func (lfe *LocalFileSystemEntry) IsDir() bool {
	return lfe.osFileInfo.IsDir()
}

func (lfe *LocalFileSystemEntry) IsFile() bool {
	return !lfe.osFileInfo.IsDir()
}


// type LocalFileSystem
// ----------------------------------------------------------------
//
type LocalFileSystem struct {
	types.FileSystem
}

// Warning:
// Use your own path mangling functionality in production code.
// This can be quite non-trivial depending on the operating system.
// The code below is not sufficient for production servers.
func rfsMangle(fpath string) (string, error) {
	// if strings.Contains(fpath, "..") {
	// 	return "<invalid>", errors.New("Invalid path")
	// }
	if len(fpath) > 0 && fpath[0] == '/' {
		fpath = fpath[1:]
	}
	// fpath = "/tmp/test-sftpd/" + fpath
	return fpath, nil
}

func isDir(osFile *os.File) bool {
	osFileStat, _ := os.Stat(osFile.Name())
	return osFileStat.IsDir()
}
// List
func (lfs *LocalFileSystem) List(fpath string) ([]LocalFileSystemEntry, error) {
	p, err := rfsMangle(fpath)
	if err != nil {
		return nil, err
	}
	osFile, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	// ファイルまたはディレクトリを読んで FsEntry 一覧を取得する
	var rs = []LocalFileSystemEntry{}
	if isDir(osFile) {
		fileInfos, _ := osFile.Readdir(0)
		rs = make([]LocalFileSystemEntry, len(fileInfos))
		for i, fileInfo := range fileInfos {
			rs[i].name = fileInfo.Name()
			rs[i].osFileInfo = fileInfo
		}
	} else {
		rs = make([]LocalFileSystemEntry, 1)
		osFileInfo, _ := os.Stat(osFile.Name())
		rs[0].name = osFileInfo.Name()
		rs[0].osFileInfo = osFileInfo
	}

	return rs, nil
}
