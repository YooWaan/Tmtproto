package fileapi

import (
	"fmt"
	"testing"
	"fileapi"
)

// TestList
func TestList(t *testing.T) {
	var fs fileapi.LocalFS
	var fses, _ = fs.List(".")

	// fmt.Printf("%v\n",fses)

	testFileName := "localfs_test.go"
	hasFile := false
	for _, fse := range fses {
		name := fse.Name()
		fmt.Printf("%v\n",name)
		if name == testFileName {
			hasFile = true
		}
	}

    if !hasFile {
        t.Errorf("%s isn't exist or error occurred during List().", testFileName)
    }
}
