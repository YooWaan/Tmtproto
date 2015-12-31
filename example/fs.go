package main

import (
	"fmt"
	"fileapi"
)

func main() {
	var fs fileapi.LocalFS
	var dirIO, _ = fs.OpenDir(".")
	var dir, _ = dirIO.Readdir(1024)

	var_dump( dir )

	fmt.Println("example/fs done.")
}

func var_dump(v ...interface{}) {
	for _, vv := range(v) {
		fmt.Printf("%#v\n", vv)
	}
}
