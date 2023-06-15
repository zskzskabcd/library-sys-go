package v1

import (
	"library-sys-go/internal/api"
	"library-sys-go/internal/middleware"
	"library-sys-go/internal/model"
	"library-sys-go/pkg/resp"

	"github.com/gin-gonic/gin"
)

// SaveReader 新增/更新读者信息 godoc
// @Summary 新增/更新读者信息
// @Description 新增/更新读者信息
// @Tags 读者
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param reader body model.Reader true "读者信息"
// @Success 200 {object} resp.Resp
// @Router /reader [post]
func SaveReader(c *gin.Context) {
	var reader model.Reader
	if err := c.ShouldBindJSON(&reader); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	// 加密密码
	if reader.Key != "" {
		reader.EncryptPassword()
	}
	err := reader.Query().Save(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.Success(c)
}

// DeleteReader 删除读者 godoc
// @Summary 删除读者
// @Description 删除读者
// @Tags 读者
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id query int true "读者ID"
// @Success 200 {object} resp.Resp
// @Router /reader [delete]
func DeleteReader(c *gin.Context) {
	var req struct {
		ID uint `json:"id" form:"id" query:"id" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	reader := model.Reader{}
	reader.ID = uint(req.ID)
	err := reader.Query().Delete(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.Success(c)
}

// ListReader 查询读者列表 godoc
// @Summary 查询读者列表
// @Description 查询读者列表
// @Tags 读者
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param keyword query string false "关键字"
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} resp.RespList[model.Reader]
// @Router /reader/list [get]
func ListReader(c *gin.Context) {
	var req struct {
		Keyword string `json:"keyword" form:"keyword"`
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

// GetReader 查询读者详情 godoc
// @Summary 查询读者详情
// @Description 查询读者详情
// @Tags 读者
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id query int true "读者ID"
// @Success 200 {object} resp.Resp{data=model.Reader}
// @Router /reader [get]
func GetReader(c *gin.Context) {
	var req struct {
		ID int `json:"id" binding:"required" example:"1" form:"id"`
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

type LoginResp struct {
	Token string `json:"token" example:"xxx"`
}

type LoginReaderReq struct {
	StudentNo string `json:"studentNo" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// 读者登陆 godoc
// @Summary 读者登陆
// @Description 读者登陆
// @Tags 读者
// @Accept json
// @Produce json
// @Param reader body LoginReaderReq true "读者登陆信息"
// @Success 200 {object} resp.Resp{data=LoginResp}
// @Router /reader/login [post]
func LoginReader(c *gin.Context) {
	var req LoginReaderReq
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
		Role:     []middleware.Role{middleware.RoleReader},
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

type UpdateReaderPasswordReq struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

// 读者修改密码 godoc
// @Summary 读者修改密码
// @Description 读者修改密码
// @Tags 读者
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param req body UpdateReaderPasswordReq true "读者修改密码信息"
// @Success 200 {object} resp.Resp
// @Router /reader/password [post]
func UpdateReaderPassword(c *gin.Context) {
	var req UpdateReaderPasswordReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	// 获取用户ID
	user := c.MustGet("user").(*middleware.UserClaims)
	reader := model.Reader{}
	err := reader.Query().Where("id = ?", user.ID).First(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	if !reader.ComparePassword(req.OldPassword) {
		resp.Error(c, resp.CodeInternalServer, "旧密码错误")
		return
	}
	reader.Key = req.NewPassword
	reader.EncryptPassword()
	err = reader.Query().Updates(&reader).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.Success(c)
}
