package v1

import (
	"backend/app/models/education"
	"backend/app/models/education_detail"
	"backend/app/models/other"
	"backend/app/models/project"
	"backend/app/models/project_detail"
	"backend/app/models/resume"
	"backend/app/models/resume_basic"
	"backend/app/models/resume_basic_field"
	"backend/app/models/work_experience"
	"backend/app/models/work_experience_detail"
	"backend/app/policies"
	"backend/app/requests"
	"backend/pkg/app"
	"backend/pkg/auth"
	"backend/pkg/database"
	"backend/pkg/paginator"
	"backend/pkg/response"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

type ResumesController struct {
	BaseAPIController
}

type RemoveIdStruct struct {
	RemoveIds []uint64 `json:"removeIds"`
}

func (ctrl *ResumesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	// response.Abort403(c, "没权")
	// return

	query := database.DB.Model(resume.Resume{}).Order("updated_at desc").Where("user_id = ?", auth.CurrentUserID(c))
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
	resumeModel := resume.Get(c.Param("id"))

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
		ModuleOrder: "resumeBasic,education,project,workExperience",
		// 创建简历基础信息关联
		ResumeBasic: &resume_basic.ResumeBasic{

			Name: resume_basic_field.ResumeBasicField{
				Value:   currentUser.Nickname,
				Visible: true,
			},
			Email: resume_basic_field.ResumeBasicField{
				Value:   currentUser.Email,
				Visible: true,
				Label:   "邮箱",
				Icon:    "mail",
			},
			Mobile: resume_basic_field.ResumeBasicField{
				Value:   currentUser.Phone,
				Visible: true,
				Label:   "电话",
				Icon:    "call",
			},
			Job: resume_basic_field.ResumeBasicField{
				Value:   "设计师",
				Visible: true,
				Label:   "求职岗位",
			},
		},
		// 创建教育经历
		Education: &education.Education{
			Label:           "教育经历",
			Visible:         true,
			ContentType:     "education2",
			ModuleTitleType: request.LayoutType,
			EducationDetails: []*education_detail.EducationDetail{
				{
					Name:             "XXXX大学",
					StartOn:          "2000-09",
					EndOn:            "2004-06",
					UniversityMajors: "XX专业",
					SortIndex:        0,
					Desc:             "这是一段关于大学的介绍。",
				},
			},
		},
		Project: &project.Project{
			Label:           "项目经历",
			Visible:         true,
			ContentType:     "project2",
			ModuleTitleType: request.LayoutType,
			ProjectDetails: []*project_detail.ProjectDetail{
				{
					Name:      "XXXX项目",
					StartOn:   "2000-09",
					EndOn:     "2004-06",
					Role:      "负责人",
					SortIndex: 0,
					Desc:      "这是一段关于项目的介绍。以及本人负责的具体内容",
				},
			},
		},
		WorkExperience: &work_experience.WorkExperience{
			Label:           "工作经历",
			Visible:         true,
			ContentType:     "workExperience2",
			ModuleTitleType: request.LayoutType,
			WorkExperienceDetails: []*work_experience_detail.WorkExperienceDetail{
				{
					Name:      "XXXX公司",
					StartOn:   "2000-09",
					EndOn:     "2004-06",
					JobTitle:  "XXX岗位",
					SortIndex: 0,
					Desc:      "这是一段关于自身在公司主要负责的工作内容",
				},
			},
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
	resumeModel := resume.Get(c.Param("id"))

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

	resumeModel.ModuleOrder = resumeModel.ModuleOrder + ",education"

	rowsAffected := resumeModel.Save()
	if rowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "添加失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) AddWorkExperience(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	if resumeModel.WorkExperience != nil {
		message := "工作经历已存在"
		response.BadRequest(c, errors.New(message), message)
		return
	}

	database.DB.Model(&resumeModel).Association("WorkExperience").Append(&work_experience.WorkExperience{
		Label:           "工作经历",
		Visible:         true,
		ContentType:     "workExperience1",
		ModuleTitleType: resumeModel.LayoutType,
	})

	resumeModel.ModuleOrder = resumeModel.ModuleOrder + ",workExperience"

	rowsAffected := resumeModel.Save()
	if rowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "添加失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) AddProject(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	if resumeModel.Project != nil {
		message := "项目经历已存在"
		response.BadRequest(c, errors.New(message), message)
		return
	}

	database.DB.Model(&resumeModel).Association("Project").Append(&project.Project{
		Label:           "项目经历",
		Visible:         true,
		ContentType:     "project1",
		ModuleTitleType: resumeModel.LayoutType,
	})

	resumeModel.ModuleOrder = resumeModel.ModuleOrder + ",project"

	rowsAffected := resumeModel.Save()
	if rowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "添加失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) AddOther(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	if len(resumeModel.Others) != 0 {
		message := "其他自定义模块已存在"
		response.BadRequest(c, errors.New(message), message)
		return
	}

	initOther := other.Other{
		Label:           "其他",
		Visible:         true,
		ContentType:     "other1",
		ModuleTitleType: resumeModel.LayoutType,
	}

	database.DB.Model(&resumeModel).Association("Others").Append(&[]other.Other{
		initOther,
	})

	resumeModel.ModuleOrder = resumeModel.ModuleOrder + ",others"

	rowsAffected := resumeModel.Save()
	if rowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "添加失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) UpdateResumeLayoutType(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := resume.Resume{}
	if err := c.ShouldBind(&request); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	resumeModel.LayoutType = request.LayoutType

	if resumeModel.Education != nil {
		resumeModel.Education.ModuleTitleType = request.LayoutType
	}

	if resumeModel.Project != nil {
		resumeModel.Project.ModuleTitleType = request.LayoutType
	}

	if resumeModel.WorkExperience != nil {
		resumeModel.WorkExperience.ModuleTitleType = request.LayoutType
	}

	if len(resumeModel.Others) != 0 {
		for _, v := range resumeModel.Others {
			v.ModuleTitleType = request.LayoutType
		}
	}

	fmt.Println(resumeModel.Others, "resumeModel.Others")

	result := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(&resumeModel)

	if result.RowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) UpdateResumeBasic(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))

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

	fmt.Println("===================================")
	fmt.Println(resumeModel)
	fmt.Println("===================================")

	result := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(&resumeModel)

	if result.RowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) UpdateEducation(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := education.Education{}
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	removeIdsRequest := RemoveIdStruct{}
	if err := c.ShouldBindBodyWith(&removeIdsRequest, binding.JSON); err != nil {
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

	removeIdsLen := len(removeIdsRequest.RemoveIds)
	if removeIdsLen != 0 {
		database.DB.Delete(&education_detail.EducationDetail{}, removeIdsRequest.RemoveIds)
	}

	if result.RowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) UpdateWorkExperience(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := work_experience.WorkExperience{}
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	removeIdsRequest := RemoveIdStruct{}
	if err := c.ShouldBindBodyWith(&removeIdsRequest, binding.JSON); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	request.ID = resumeModel.ResumeBasic.ID
	resumeModel.WorkExperience = &request

	result := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(&resumeModel)

	removeIdsLen := len(removeIdsRequest.RemoveIds)
	if removeIdsLen != 0 {
		database.DB.Delete(&work_experience_detail.WorkExperienceDetail{}, removeIdsRequest.RemoveIds)
	}

	if result.RowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) UpdateProject(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := project.Project{}
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	removeIdsRequest := RemoveIdStruct{}
	if err := c.ShouldBindBodyWith(&removeIdsRequest, binding.JSON); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	request.ID = resumeModel.ResumeBasic.ID
	resumeModel.Project = &request

	result := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(&resumeModel)

	removeIdsLen := len(removeIdsRequest.RemoveIds)
	if removeIdsLen != 0 {
		database.DB.Delete(&project_detail.ProjectDetail{}, removeIdsRequest.RemoveIds)
	}

	if result.RowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ResumesController) UpdateOthers(c *gin.Context) {
	resumeModel := resume.Get(c.Param("id"))

	if resumeModel.ID == 0 {
		response.Abort404(c)
		return
	}

	type Request struct {
		Others    []*other.Other `json:"others"`
		RemoveIds []interface{}  `json:"removeIds"`
	}

	request := Request{}
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return
	}

	if ok := policies.CanModifyResume(c, resumeModel); !ok {
		response.Abort403(c)
		return
	}

	resumeModel.Others = request.Others

	result := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).
		Updates(&resumeModel)

	removeIdsLen := len(request.RemoveIds)
	if removeIdsLen != 0 {
		database.DB.Delete(&other.Other{}, request.RemoveIds)
	}

	if result.RowsAffected > 0 {
		response.Data(c, resumeModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}
