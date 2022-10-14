package v1

import (
	"errors"
	"gen-resume/app/models/education"
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
	"gorm.io/gorm"
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
	database.DB.Model(&resume.Resume{}).
		Preload(clause.Associations).
		Where("slug = ?", c.Param("slug")).
		First(&resumeModel)

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
		Name:        request.Name,
		Slug:        request.Slug,
		LayoutType:  request.LayoutType,
		UserID:      auth.CurrentUserID(c),
		ModuleOrder: "resumeBasic",
		// 创建简历基础信息关联
		ResumeBasic: &resume_basic.ResumeBasic{
			Name:        currentUser.Nickname,
			NameConfig:  "{\"visible\":true}",
			Email:       currentUser.Email,
			EmailConfig: "{\"visible\":true,\"icon\":\"mail\"}",
			Mobile:      currentUser.Phone,
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
	var resumeModel resume.Resume
	database.DB.Model(&resume.Resume{}).
		Preload(clause.Associations).
		Where("slug = ?", c.Param("slug")).
		First(&resumeModel)

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	request := resume.Resume{}
	if err := c.ShouldBind(&request); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	result := database.DB.Model(&resumeModel).Updates(request)

	if result.RowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) Delete(c *gin.Context) {
	resumeModel := resume.Get(c.Param("slug"))
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

func (ctrl *ResumesController) AddEducation(c *gin.Context) {
	var resumeModel resume.Resume
	database.DB.Model(&resume.Resume{}).
		Preload(clause.Associations).
		Where("slug = ?", c.Param("slug")).
		First(&resumeModel)

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	if resumeModel.Education != nil {
		message := "教育经历已存在"
		response.BadRequest(c, errors.New(message), message)
		return
	}

	database.DB.Model(&resumeModel).Association("Education").Append(&education.Education{
		Label:           "教育经历",
		Visible:         true,
		ContentType:     "education1",
		ModuleTitleType: resumeModel.LayoutType,
	})

	resumeModel.ModuleOrder = resumeModel.ModuleOrder + ",educations"

	rowsAffected := resumeModel.Save()
	if rowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "添加失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) UpdateResumeBasic(c *gin.Context) {
	var resumeModel resume.Resume
	database.DB.Model(&resume.Resume{}).
		Preload(clause.Associations).
		Where("slug = ?", c.Param("slug")).
		First(&resumeModel)

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := resume_basic.ResumeBasic{}
	if err := c.ShouldBind(&request); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	request.ID = resumeModel.ResumeBasic.ID
	resumeModel.ResumeBasic = &request

	result := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(&resumeModel)

	if result.RowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) UpdateEducation(c *gin.Context) {
	var resumeModel resume.Resume
	database.DB.Model(&resume.Resume{}).
		Preload(clause.Associations).
		Where("slug = ?", c.Param("slug")).
		First(&resumeModel)

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := education.Education{}
	if err := c.ShouldBind(&request); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	request.ID = resumeModel.ResumeBasic.ID
	resumeModel.Education = &request

	result := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(&resumeModel)

	if result.RowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}
