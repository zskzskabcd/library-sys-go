package v1

import (
	"library-sys-go/internal/middleware"
	"library-sys-go/internal/model"
	"library-sys-go/pkg/resp"
	"log"

	"github.com/gin-gonic/gin"
)

type AdminLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 管理员登录 godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员
// @Accept json
// @Produce json
// @Param userInfo body AdminLoginReq true "管理员登录信息"
// @Success 200 {object} resp.Resp{data=LoginResp}
// @Router /admin/login [post]
func AdminLogin(c *gin.Context) {
	var req AdminLoginReq
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	admin := model.Admin{}
	err := admin.Query().Where("username = ?", req.Username).First(&admin).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, "用户名或密码错误")
		return
	}
	if !admin.ComparePassword(req.Password) {
		resp.Error(c, resp.CodeInternalServer, "用户名或密码错误")
		return
	}
	// 生成token
	token, err := middleware.GenerateToken(middleware.UserClaims{
		ID:       admin.ID,
		UserName: admin.Name,
		Role:     []middleware.Role{middleware.RoleAdmin},
		UserType: "admin",
	})
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.SuccessData(c, gin.H{
		"token": token,
	})
}

// 管理员修改密码 godoc
// @Summary 管理员修改密码
// @Description 管理员修改密码
// @Tags 管理员
// @Accept json
// @Produce json
// @Param oldPassword query string true "旧密码"
// @Param newPassword query string true "新密码"
// @Success 200 {object} resp.Resp
// @Router /admin/changePassword [post]
func AdminChangePassword(c *gin.Context) {
	var req struct {
		OldPassword string `json:"oldPassword" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		resp.Error(c, resp.CodeParamsInvalid, err.Error())
		return
	}
	admin := model.Admin{}
	err := admin.Query().Where("id = ?", c.GetInt64("id")).First(&admin).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	if !admin.ComparePassword(req.OldPassword) {
		resp.Error(c, resp.CodeInternalServer, "旧密码错误")
		return
	}
	admin.Password = req.NewPassword
	admin.EncryptPassword()
	err = admin.Query().Updates(&admin).Error
	if err != nil {
		resp.Error(c, resp.CodeInternalServer, err.Error())
		return
	}
	resp.Success(c)
}

// 创建管理员 仅限控制台
func AddAdmin() {
	// 是否存在管理员
	admin := model.Admin{}
	err := admin.Query().First(&admin).Error
	if err == nil {
		log.Println("已存在管理员")
		return
	}
	admin = model.Admin{
		Name:     "admin",
		Password: "123456",
	}
	admin.EncryptPassword()
	err = admin.Query().Create(&admin).Error
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("创建管理员成功")
}
