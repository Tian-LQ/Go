package gin_v3

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context HTTP 请求上下文
type Context struct {
	// 原生的上下文对象
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求相关信息
	Path   string
	Method string
	Params map[string]string
	// 响应相关信息
	StatusCode int
}

// newContext 创建一个HTTP请求上下文对象
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// Param 获取当前请求URL当中 param [key,value] 数据
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm 获取请求报文 form 表单数据当中的指定 [key,value]
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 获取请求 URL 查询参数当中的指定 [key,value]
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置响应报文当中的状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 向响应报文当中的响应首部添加 [key,value] 值对
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 字符串类型数据写入响应报文 Body
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON Json 类型数据写入响应报文 Body
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		panic(err)
	}
}

// Data 二进制字符流类型数据写入响应报文 Body
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML Html 类型数据写入响应报文 Body
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
