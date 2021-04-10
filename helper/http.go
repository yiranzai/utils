package helper

import (
	"github.com/gin-gonic/gin"
)

// 是否是Ajax请求
func IsAjax(c *gin.Context) bool {
	return "XMLHttpRequest" == c.GetHeader("X-Requested-With")
}

// 获取完整的URL
func FullUrl(c *gin.Context) string {
	return "http://" + c.Request.Host + c.Request.URL.Path
}
