package wire

import (
	"github.com/google/wire"
)

func NewMySQLClient(url string) *MySQLClient {
	return &MySQLClient{url: url}
}

type MySQLClient struct {
	url string
}

func (c *MySQLClient) Exec(query string, args ...interface{}) string {
	return "data"
}

func NewApp(client *MySQLClient) *App {
	return &App{client: client}
}

type App struct {
	// App 持有唯一的 MySQLClient 实例
	client *MySQLClient
}

func (a *App) GetData(query string, args ...interface{}) string {
	data := a.client.Exec(query, args...)
	return data
}

func Handle(url string) *App {
	// 使用 Wire 获取 ServiceB 实例
	wire.Build(NewMySQLClient, NewApp)
	return nil
}
