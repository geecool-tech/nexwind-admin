package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nexwind.net/admin/server/constant"
)

type Resp struct {
	CODE_FAIL            int
	CODE_SUCCESS         int
	CODE_TOKEN_NOT_EXIST int
	CODE_UN_AUTH         int
}

func InitResp() Resp {
	return Resp{
		CODE_FAIL:            constant.CODE_FAIL,
		CODE_SUCCESS:         constant.CODE_SUCCESS,
		CODE_TOKEN_NOT_EXIST: constant.CODE_TOKEN_NOT_EXIST,
		CODE_UN_AUTH:         constant.CODE_UN_AUTH,
	}
}

// SuccessRes success common response.
func (Resp) SuccessRes(c *gin.Context, data any) {
	c.JSON(http.StatusOK, SuccessRes{
		Code:   constant.CODE_SUCCESS,
		Status: "ok",
		Msg:    GetMsg(constant.CODE_SUCCESS),
		Data:   data,
	})
}

// ErrorRes error common response.
func (Resp) ErrorRes(c *gin.Context, res ErrorRes) {
	code := res.Code

	if code == 0 {
		code = -1
	}
	if res.Msg == "" {
		res.Msg = GetMsg(code)
	}
	c.JSON(http.StatusOK, ErrorRes{
		Code: code,
		Msg:  res.Msg,
	})
}

// PageRes pagination common response.
func (Resp) PageRes(c *gin.Context, list ListRes) {
	c.JSON(http.StatusOK, SuccessRes{
		Code: constant.CODE_SUCCESS,
		Msg:  GetMsg(constant.CODE_SUCCESS),
		Data: list,
	})
}
