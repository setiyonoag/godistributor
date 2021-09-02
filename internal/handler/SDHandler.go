package handler

import (
	"github.com/gin-gonic/gin"
	"godistributor/internal/entity"
	"godistributor/internal/service"
	"godistributor/pkg/response"
	"net/http"
	"strconv"
)

type SDHandler struct {
	sDService service.InterfaceSDService
}

func NewSDHandler(sDService service.InterfaceSDService) *SDHandler {
	var sDHandler = SDHandler{}
	sDHandler.sDService = sDService
	return &sDHandler
}

func (h *SDHandler) GetSD(c *gin.Context) {
	result, err := h.sDService.GetSD()
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}
	if result == nil {
		result = []entity.SDModelView{}
	}

	response.ResponseOkWithData(c, result)
}

func (h *SDHandler) GetOneSD(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id_sd"))
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
	}

	result, err := h.sDService.GetOneSD(id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
	}

	response.ResponseOkWithData(c, result)

}