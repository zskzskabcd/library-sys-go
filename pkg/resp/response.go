package resp

import (
	"encoding/json"
	"library-sys-go/pkg/encode"

	"github.com/gin-gonic/gin"
)

// Code
const (
	CodeSuccess        = 200
	CodeParamsInvalid  = 400
	CodeUnauthorized   = 401
	CodeForbidden      = 403
	CodeNotFound       = 404
	CodeInternalServer = 500
)

// Success 操作成功返回
func Success(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": CodeSuccess,
		"msg":  "success",
	})
}

// SuccessData 操作成功返回
func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": CodeSuccess,
		"msg":  "success",
		"data": data,
	})
}

// SuccessList 操作成功返回
func SuccessList(c *gin.Context, data interface{}, total int64) {
	c.JSON(200, gin.H{
		"code":  CodeSuccess,
		"msg":   "success",
		"data":  data,
		"count": total,
	})
}

// SuccessDataEncrypt 操作成功返回 加密
func SuccessDataEncrypt(c *gin.Context, data interface{}) {
	dataJson, err := json.Marshal(data)
	if err != nil {
		Error(c, CodeInternalServer, "数据加密失败")
		return
	}
	url := c.Request.URL.String()
	dataEncrypted := encode.Encode(url, string(dataJson))
	// base64编码
	//dataBase64 := base64.StdEncoding.EncodeToString(dataEncrypted)
	c.JSON(200, gin.H{
		"code":    CodeSuccess,
		"msg":     "success",
		"data":    dataEncrypted,
		"encrypt": true,
	})
}

// Error 操作失败返回
func Error(c *gin.Context, code int, msg string) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}
