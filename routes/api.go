// Package routes 注册路由
package routes

import (
	apiV1 "gen-resume/app/controllers/api/v1"
	"gen-resume/app/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/api/v1")
	{
		usersGroup := v1.Group("/users")
		{
			users := new(apiV1.UsersController)
			// 判断手机是否存在
			usersGroup.POST("/phone-exist", users.IsPhoneExist)
			// 判断邮箱是否存在
			usersGroup.POST("/email-exist", users.IsEmailExist)
		}

		captchaGroup := v1.Group("/captcha")
		{
			captcha := new(apiV1.CaptchaController)
			// 图片验证码，需要添加限流
			captchaGroup.GET("", captcha.ShowCaptcha)
		}

		authGroup := v1.Group("/auth")
		verifyCode := new(auth.VerifyCodeController)
		signUpController := new(auth.SignUpController)

		authGroup.POST("/verify-codes/phone", verifyCode.SendUsingPhone)
		authGroup.POST("/verify-codes/email", verifyCode.SendUsingEmail)
		authGroup.POST("/sign-up/using-phone", signUpController.SignUpUsingPhone)
	}
}
