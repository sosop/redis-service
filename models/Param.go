package models

import "strings"

const (
	SUC = iota
	PARAM_ERROR
	SERVER_ERROR
)

type (
	Param struct {
		Db   int
		Cmd  string
		Args []string
	}
	Ret struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)

func NewSucRet(data interface{}) Ret {
	return Ret{SUC, "success", data}
}

func NewServerRet(msg string) Ret {
	return Ret{Code: SERVER_ERROR, Msg: msg}
}

func NewParamRet() Ret {
	return Ret{Code: PARAM_ERROR, Msg: "参数不正确"}
}

func (p *Param) valid() bool {
	if strings.TrimSpace(p.Cmd) == "" {
		return false
	}
	return true
}
