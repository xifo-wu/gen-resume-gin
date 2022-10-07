package v1

import (
	"gen-resume/app/models/resume"
	"gen-resume/app/models/resume_basic"
	"gen-resume/app/policies"
	"gen-resume/app/requests"
	"gen-resume/pkg/app"
	"gen-resume/pkg/auth"
	"gen-resume/pkg/database"
	"gen-resume/pkg/paginator"
	"gen-resume/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type ResumesController struct {
	BaseAPIController
}

func (ctrl *ResumesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	query := database.DB.Model(resume.Resume{}).Where("user_id = ?", auth.CurrentUserID(c))
	var resumes []resume.Resume

	pager := paginator.Paginate(
		c,
		query,
		&resumes,
		app.V1URL(database.TableName(&resume.Resume{})),
		10,
	)

	response.JSON(c, gin.H{
		"data": resumes,
		"meta": pager,
	})
}

func (ctrl *ResumesController) Show(c *gin.Context) {
	var resumeModel resume.Resume
	database.DB.Model(&resume.Resume{}).Preload(clause.Associations).Where("slug = ?", c.Param("id")).First(&resumeModel)

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	response.Data(c, resumeModel)
}

func (ctrl *ResumesController) Store(c *gin.Context) {
	request := requests.ResumeRequest{}
	if ok := requests.Validate(c, &request, requests.ResumeSave); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)

	resumeModel := resume.Resume{
		Name:       request.Name,
		Slug:       request.Slug,
		LayoutType: request.LayoutType,
		UserID:     auth.CurrentUserID(c),
		// 创建简历基础信息关联
		ResumeBasic: resume_basic.ResumeBasic{
			Name:   currentUser.Nickname,
			Email:  currentUser.Email,
			Mobile: currentUser.Phone,
		},
	}

	resumeModel.Create()
	if resumeModel.ID > 0 {
		response.Created(c, resumeModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) Update(c *gin.Context) {

	resumeModel := resume.Get(c.Param("id"))
	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.ResumeRequest{}
	if ok := requests.Validate(c, &request, requests.ResumeSave); !ok {
		return
	}

	resumeModel.Name = request.Name
	resumeModel.Slug = request.Slug
	resumeModel.LayoutType = request.LayoutType

	rowsAffected := resumeModel.Save()
	if rowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) Delete(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))
	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := resumeModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
