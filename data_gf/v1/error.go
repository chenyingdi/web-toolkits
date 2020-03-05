package data

import "errors"

var (
	NotFound = errors.New("资源不存在！")
	GetTypeError = errors.New("请输入正确的查询类型！")
)
