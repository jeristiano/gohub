package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器

type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupPhoneExistsRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}
	//  检查数据库并返回响应

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})

}

// SignupUsingPhone 使用手机和验证码进行注册

func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	// 1. 验证表单
	// 初始化请求对象
	request := requests.SignupUsingPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}
	// 2. 验证成功，创建数据
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()
	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
