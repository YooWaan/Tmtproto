package server

import (
	"types"
)


// 実際に処理をするもの暫定的に struct にする function にするかも
type Flow struct {
	// 認証 Callback
	authCallback types.AuthCallback
	// TODO file API
}

// 
type FlowBuilder struct {}

// 認証用 Callback を追加
func (fb *FlowBuilder) auth(authtype string) (*FlowBuilder) {
	// 文字列をみて callback を生成する
	return fb
}

// File API を追加
func (fb *FlowBuilder) fileapi(fileapitype string) (*FlowBuilder) {
	// 文字列をみて fileapi を生成する

	return fb
}

func (fb *FlowBuilder) build() (*Flow) {
	return &Flow{authCallback: nil}
}
