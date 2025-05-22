package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ragnarok/internal/dto"
	"ragnarok/internal/interface"
	"strconv"
)

type BrimoExpeditionController struct {
	brimoSvc _interface.BrimoExpeditionService
}

func NewBrimoExpeditionController(svc _interface.BrimoExpeditionService) *BrimoExpeditionController {
	return &BrimoExpeditionController{
		brimoSvc: svc,
	}
}

func (ctrl *BrimoExpeditionController) GetListBrimoExpedition(c *gin.Context) {
	payload := dto.BrimoExpeditionRequest{}
	err := c.ShouldBindQuery(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			ResponseCode:    "01",
			ResponseMessage: err.Error(),
			ResponseData:    nil,
		})
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	data, totalRecords, err := ctrl.brimoSvc.GetListBrimoExpedition(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			ResponseCode:    "01",
			ResponseMessage: err.Error(),
			ResponseData:    nil,
		})
	}

	limit := 10
	pagination := dto.SetupPagination(data, page, limit, totalRecords)
	response := dto.Response{
		ResponseCode:    "00",
		ResponseMessage: "Success",
		ResponseData:    pagination,
	}

	c.JSON(http.StatusOK, response)
}
