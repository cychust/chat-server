package controller

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}
type Msg struct {
	Msg  string      `json:"msg"`
	Code int         `json:"innerCode,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (c *BaseController) RespMsg(code int, msg string, data interface{}) {
	c.Ctx.ResponseWriter.Status = code
	c.Data["json"] = Msg{
		Msg:  msg,
		Data: data,
	}
	c.ServeJSON()
}

func (c *BaseController) ErrorMsg(resCode int, msg string, innerCode int) {
	c.Ctx.ResponseWriter.Status = resCode
	c.Data["json"] = Msg{
		Msg:  msg,
		Code: innerCode,
	}
	c.ServeJSON()
}
