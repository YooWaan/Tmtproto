package types

import (
	"testing"
	"os"
	"io/ioutil"
)

func TestInit(t *testing.T) {
	log := NewAppLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	//"fmt"
	//fmt.Printf("logger %s", log)

	log.Trace.Printf("trace")
	log.Info.Printf("info")
	log.Warn.Printf("warn")
	log.Err.Printf("err")
}
