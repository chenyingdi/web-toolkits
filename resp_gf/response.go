package resp

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 操作成功
func (res *Response) Success(r *ghttp.Request, data interface{}) {
	res.Code = http.StatusOK
	res.Msg = "操作成功！"
	res.Data = data

	r.Response.WriteJsonExit(res)
}

// 无法连接服务器
func (res *Response) BadGateway(r *ghttp.Request) {
	res.Code = http.StatusBadGateway
	res.Msg = "无法连接服务器！"

	r.Response.WriteJsonExit(res)
}

// 服务器内部错误
func (res *Response) ServerError(r *ghttp.Request, data interface{}) {
	res.Code = http.StatusInternalServerError
	res.Msg = "服务器内部错误！"
	res.Data = data

	r.Response.WriteJsonExit(res)
}

// 账户或密码错误
func (res *Response) CheckPasswordFailed(r *ghttp.Request) {
	res.Code = http.StatusForbidden
	res.Msg = "账户或密码错误！"

	r.Response.WriteJsonExit(res)
}

// 资源不存在
func (res *Response) NotFound(r *ghttp.Request) {
	res.Code = http.StatusNotFound
	res.Msg = "资源不存在！"

	r.Response.WriteJsonExit(res)
}

// 错误的请求
func (res *Response) BadRequest(r *ghttp.Request) {
	res.Code = http.StatusBadRequest
	res.Msg = "错误的请求！"

	r.Response.WriteJsonExit(res)
}

// 密码错误
func (res *Response) PasswordError(r *ghttp.Request) {
	res.Code = http.StatusUnauthorized
	res.Msg = "密码错误！"

	r.Response.WriteJsonExit(res)
}

// 密码一样
func (res *Response) TheSamePassword(r *ghttp.Request) {
	res.Code = http.StatusConflict
	res.Msg = "密码一致"

	r.Response.WriteJsonExit(res)
}

// 未授权
func (res *Response) UnAuthorized(r *ghttp.Request) {
	res.Code = http.StatusForbidden
	res.Msg = "未授权"

	r.Response.WriteJsonExit(res)
}
