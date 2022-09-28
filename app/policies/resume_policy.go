// Package policies 用户授权
package policies

import (
	"gen-resume/app/models/resume"
	"gen-resume/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyResume(c *gin.Context, _resume resume.Resume) bool {
	return auth.CurrentUserID(c) == _resume.UserID
}
