package v1

import (
	"library-sys-go/internal/api"
	"library-sys-go/internal/middleware"
	"library-sys-go/internal/model"
	"library-sys-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

// SaveReader 新增/更新读者信息
func SaveReader(c *gin.Context) {
	var reader model.Reader
	if err := c.ShouldBindJSON(&reader); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	err := reader.Query().Save(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.Success(c)
}

// DeleteReader 删除读者
func DeleteReader(c *gin.Context) {
	var req struct {
		ID int `json:"id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	reader := model.Reader{}
	err := reader.Query().Delete(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.Success(c)
}

// ListReader 查询读者列表
func ListReader(c *gin.Context) {
	var req struct {
		Keyword string `json:"keyword"`
		api.Pagination
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	reader := model.Reader{}
	readersQuery := reader.Query()
	if req.Keyword != "" {
		readersQuery = readersQuery.Where("name LIKE ?", "%"+req.Keyword+"%")
	}
	var readers []model.Reader
	var total int64
	err := readersQuery.Count(&total).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	err = readersQuery.Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&readers).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 去除敏感信息
	for i := range readers {
		readers[i].Key = ""
	}
	resp.SuccessList(c, readers, total)
}

// GetReader 查询读者详情
func GetReader(c *gin.Context) {
	var req struct {
		ID int `json:"id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	reader := model.Reader{}
	err := reader.Query().Where("id = ?", req.ID).First(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	// 去除敏感信息
	reader.Key = ""
	resp.SuccessData(c, reader)
}

// 读者登陆
func LoginReader(c *gin.Context) {
	var req struct {
		StudentNo string `json:"studentNo" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	reader := model.Reader{}
	err := reader.Query().Where("student_no = ?", req.StudentNo).First(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, "用户名或密码错误")
		return
	}
	if !reader.ComparePassword(req.Password) {
		resp.Error(c, resp.CodeInternalServer, "用户名或密码错误")
		return
	}
	// 去除敏感信息
	reader.Key = ""
	// 生成JWT
	token, err := middleware.GenerateToken(middleware.UserClaims{
		ID:       reader.ID,
		UserName: reader.Name,
		Role:     "reader",
		UserType: "reader",
	})
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.SuccessData(c, gin.H{
		"token": token,
	})
}
