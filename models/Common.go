package models

import (
	"github.com/ziyoubiancheng/xcms/consts"
)

// JsonResult 用于返回ajax请求的基类
type JsonResult struct {
	Code consts.JsonResultCode `json:"code"`
	Msg  string                `json:"msg"`
	Data interface{}           `json:"data"`
}

type ListJsonResult struct {
	Code  consts.JsonResultCode `json:"code"`
	Msg   string                `json:"msg"`
	Data  interface{}           `json:"data"`
	Count int                   `json:count`
}
