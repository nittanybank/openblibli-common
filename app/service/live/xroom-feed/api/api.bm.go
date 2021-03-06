// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package api is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api.proto
*/
package api

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathRecPoolGetList = "/live.xroomfeed.v1.RecPool/GetList"

// =================
// RecPool Interface
// =================

type RecPoolBMServer interface {
	// 根据模块位置获取投放列表 position=>RoomItem
	GetList(ctx context.Context, req *RecPoolReq) (resp *RecPoolResp, err error)
}

var RecPoolSvc RecPoolBMServer

func recPoolGetList(c *bm.Context) {
	p := new(RecPoolReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := RecPoolSvc.GetList(c, p)
	c.JSON(resp, err)
}

// RegisterRecPoolBMServer Register the blademaster route
func RegisterRecPoolBMServer(e *bm.Engine, server RecPoolBMServer) {
	RecPoolSvc = server
	e.GET("/live.xroomfeed.v1.RecPool/GetList", recPoolGetList)
}
