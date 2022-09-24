package v1

import (
	"gen-resume/pkg/captcha"
	"gen-resume/pkg/logger"
	"net/http"

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
	// 返回给用户
	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
