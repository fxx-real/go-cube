package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Servicewall/go-cube/config"
)

var defaultHandler *Handler

// Init 初始化全局默认 Handler，使用内置模型。
// hosts/database/username/password 直接对应 ClickHouse 连接参数。
func Init(hosts []string, database, username, password string) {
	h, err := New(&config.ClickHouseConfig{
		Hosts:    hosts,
		Database: database,
		Username: username,
		Password: password,
	})
	if err != nil {
		panic(fmt.Sprintf("go-cube: Init failed: %v", err))
	}
	defaultHandler = h
}

// Load 使用全局 Handler 执行查询，query 为 JSON 字符串。
// 必须先调用 Init。
func Load(ctx context.Context, query string) (*QueryResponse, error) {
	if defaultHandler == nil {
		return nil, fmt.Errorf("go-cube: call Init before Load")
	}
	return defaultHandler.Load(ctx, []byte(query))
}

// HTTPHandler 返回全局 Handler 的 http.Handler，可挂载到任意路由器。
// 必须先调用 Init。
func HTTPHandler() http.Handler {
	if defaultHandler == nil {
		panic("go-cube: call Init before HTTPHandler")
	}
	return http.HandlerFunc(defaultHandler.HandleLoad)
}
