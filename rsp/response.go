package rsp

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yiranzai/utils/e"
)

// Response 返回格式
type Response struct {
	Code int         `json:"code"`
	Info string      `json:"info"`
	Data interface{} `json:"data"`
}

// Success 成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: e.Success.Code(),
		Info: e.Success.Message(),
		Data: data,
	})
}

// Error 错误
func Error(c *gin.Context, err error, data interface{}) {
	// 解析获取对应的内容
	dec := e.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code: dec.Code(),    // 错误码
		Info: dec.Message(), // 错误描述
		Data: data,
	})
}

// Download 下载文件，传入文件名和二级制内容
func Download(c *gin.Context, fileName string, buffer []byte) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s;filename*=UTF-8%s", fileName, fileName))
	c.Header("Content-Type", "application/octet-stream")
	_, _ = c.Writer.Write(buffer)
}
