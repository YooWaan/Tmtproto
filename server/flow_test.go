package server

import (
	"testing"
)


func TestBuild(t *testing.T) {
	var fb FlowBuilder = FlowBuilder{}
	var f = fb.build()
	if f == nil {
		t.Error("cant build flow")
	} else {
		t.Log("create flow")
	}
}
