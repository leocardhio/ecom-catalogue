package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leocardhio/ecom-catalogue/db/model"
	"github.com/leocardhio/ecom-catalogue/service"
	"github.com/leocardhio/ecom-catalogue/util"
)

type ITagsController interface {
	CreateTag(c *gin.Context)
	DeleteTag(c *gin.Context)
	UpdateTag(c *gin.Context)
}

type tagsController struct {
	tagsService service.ITagsService
}

func NewTagsController(service service.ITagsService) ITagsController {
	return &tagsController{
		tagsService: service,
	}
}

func (controller *tagsController) CreateTag(c *gin.Context) {
	var req model.CreateTagRequest
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	res, err := controller.tagsService.CreateTag(c, req)
	if err != nil {
		util.ResponseInternalServerError(c, err)
		return
	}

	util.ResponseCreated(c, res)
}

func (controller *tagsController) DeleteTag(c *gin.Context) {
	var req model.DeleteTagRequest
	var err error

	if err = c.ShouldBindUri(&req); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	res, err := controller.tagsService.DeleteTag(c, req)
	if err != nil {
		util.ResponseInternalServerError(c, err)
		return
	}

	util.ResponseDeleted(c, res)
}

func (controller *tagsController) UpdateTag(c *gin.Context) {
	var req model.UpdateTagRequest
	var err error

	if err = c.ShouldBindUri(&req.UpdateTagRequestURI); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	if err = c.ShouldBindJSON(&req.UpdateTagRequestBody); err != nil {
		util.ResponseBadRequest(c, err)
		return
	}

	res, err := controller.tagsService.UpdateTag(c, req)
	if err != nil {
		util.ResponseInternalServerError(c, err)
		return
	}

	util.ResponseUpdated(c, res)
}
