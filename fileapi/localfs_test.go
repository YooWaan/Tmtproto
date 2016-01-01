package fileapi

import (
	"fmt"
	"testing"
	"fileapi"
)

// TestListByFile
func TestListByFile(t *testing.T) {
	localFileSystem := fileapi.LocalFileSystem{}
	var lfsEntries, err = localFileSystem.List("../README.md")
	if err != nil {
		t.Errorf("%v",err)
		return
	}

	fmt.Printf("%v\n",lfsEntries)

	testFileName := "README.md"
	hasFile := false
	for _, lfsEntry := range lfsEntries {
		name := lfsEntry.Name()
		fmt.Printf("%v (IsDir(): %v)\n", name, lfsEntry.IsDir())
		if name == testFileName {
			hasFile = true
		}
	}

	if !hasFile {
		t.Errorf("%s isn't exist or error occurred during List().", testFileName)
	}
}

// TestListByFolder
func TestListByFolder(t *testing.T) {
	localFileSystem := fileapi.LocalFileSystem{}
	var lfsEntries, err = localFileSystem.List("../")
	if err != nil {
		t.Errorf("%v",err)
		return
	}

	fmt.Printf("%v\n",lfsEntries)

	testFileName := ".gitignore"
	hasFile := false
	for _, lfsEntry := range lfsEntries {
		name := lfsEntry.Name()
		fmt.Printf("%v (IsDir(): %v)\n", name, lfsEntry.IsDir())
		if name == testFileName {
			hasFile = true
		}
	}

	if !hasFile {
		t.Errorf("%s isn't exist or error occurred during List().", testFileName)
	}
}
