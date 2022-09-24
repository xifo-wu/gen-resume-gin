package v1

import (
	"gen-resume/pkg/captcha"
	"gen-resume/pkg/logger"
	"gen-resume/pkg/response"

	"github.com/gin-gonic/gin"
)

type CaptchaController struct {
	BaseAPIController
}

func (controller *CaptchaController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)

	response.Data(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
