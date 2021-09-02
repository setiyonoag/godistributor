package handler

import (
	"github.com/gin-gonic/gin"
	"godistributor/internal/entity"
	"godistributor/internal/service"
	"godistributor/pkg/response"
	"net/http"
	"strconv"
)

type MDHandler struct {
	mDService service.InterfaceMDService
}

func NewMDHandler(mDService service.InterfaceMDService) *MDHandler {
	var mDHandler = MDHandler{}
	mDHandler.mDService = mDService
	return &mDHandler
}

func (h *MDHandler) GetMD(c *gin.Context) {
	result, err := h.mDService.GetMD()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	if result == nil {
		result = []entity.MDModelView{}
	}

	response.ResponseOkWithData(c, result)
}

func (h *MDHandler) GetOneMD(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id_md"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	}

	result, err := h.mDService.GetOneMD(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseOkWithData(c, result)

}