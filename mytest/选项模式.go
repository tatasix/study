package mytest

import (
	"fmt"
	"net/http"
	"time"
)

// Client 是 http client 的封装
type Client struct {
	httpClient *http.Client
}

// Option 是一个函数类型，用于设置 Client 的配置选项
type Option func(*Client)

// WithTimeout 用于设置 http 请求的超时时间
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// NewClient 用于创建默认配置的 http client
func NewClient(opts ...Option) *Client {
	defaultClient := &http.Client{
		Timeout: time.Second * 5,
	}
	client := &Client{
		httpClient: defaultClient,
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func Handle() {
	// 创建一个默认 5 秒超时的 http client
	client := NewClient(WithTimeout(5 * time.Second))

	// 发送 GET 请求
	resp, err := client.httpClient.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	defer resp.Body.Close()

	// 输出 response body
	buf := make([]byte, 1024)
	n, err := resp.Body.Read(buf)
	if err != nil {
		fmt.Println("read response body failed:", err)
		return
	}
	fmt.Println(string(buf[:n]))
}
