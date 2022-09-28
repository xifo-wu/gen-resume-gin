// Package routes 注册路由
package routes

import (
	apiV1 "gen-resume/app/controllers/api/v1"
	"gen-resume/app/controllers/api/v1/auth"
	"gen-resume/app/controllers/api/v1/manager"
	"gen-resume/app/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/api/v1")
	{
		authorized := v1.Group("")
		authorized.Use(middlewares.AuthJWT())

		publicUserGroup := v1.Group("/users")
		authorizedUserGroup := authorized.Group("/users")

		{
			users := new(apiV1.UsersController)
			// 判断手机是否存在
			publicUserGroup.POST("/phone-exist", users.IsPhoneExist)
			// 判断邮箱是否存在
			publicUserGroup.POST("/email-exist", users.IsEmailExist)
			// 当前用户
			authorizedUserGroup.GET("/current", users.CurrentUser)
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
		loginController := new(auth.LoginController)
		passwordController := new(auth.PasswordController)

		authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("8-D"), verifyCode.SendUsingPhone)
		authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("8-D"), verifyCode.SendUsingEmail)
		authGroup.POST("/sign-up/using-phone", middlewares.LimitPerRoute("8-D"), signUpController.SignUpUsingPhone)
		authGroup.POST("/sign-up/using-email", middlewares.LimitPerRoute("8-D"), signUpController.SignUpUsingEmail)
		authGroup.POST("/login/using-phone", loginController.LoginByPhone)
		authGroup.POST("/login/using-password", middlewares.LimitPerRoute("10-M"), loginController.LoginByPassword)
		authGroup.POST("/login/refresh-token", loginController.RefreshToken)
		authGroup.POST("/password-reset/using-phone", middlewares.LimitPerRoute("8-D"), passwordController.ResetByPhone)
		authGroup.POST("/password-reset/using-email", middlewares.LimitPerRoute("8-D"), passwordController.ResetByEmail)

		// 客户端接口

		// 简历接口
		resumeController := new(apiV1.ResumesController)
		resumeGroup := authorized.Group("/resumes")
		{
			resumeGroup.POST("", resumeController.Store)
			resumeGroup.GET("", resumeController.Index)
			resumeGroup.GET("/:id", resumeController.Show)
			resumeGroup.PUT("/:id", resumeController.Update)
			resumeGroup.DELETE("/:id", resumeController.Delete)
		}

		// Manager API
		managerGroup := v1.Group("/manager")
		managerGroup.Use(middlewares.AuthJWT())
		{
			// 用户管理
			managerUsersController := new(manager.UsersController)
			managerUsersGroup := managerGroup.Group("/users")
			{
				managerUsersGroup.GET("", managerUsersController.Index)
			}
		}
	}
}
