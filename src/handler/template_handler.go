package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
	"github.com/sahidhossen/synmail/src/middleware"
	"github.com/sahidhossen/synmail/src/models"
)

func (h *GinHandler) CreateTemplate(c *gin.Context) {
	template := &models.Template{}

	if err := c.Bind(template); err != nil {
		log.Err(err).Msg("Required field")
		ResponseWithError(c, http.StatusBadRequest, "Required field missing!")
		return
	}

	result, err := h.DBService.CreateTemplate(template)
	if err != nil {
		log.Err(err).Msg("Template create error")
		ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, http.StatusOK, "success", "Template created!", result)
}

func (h *GinHandler) UpdateTemplate(c *gin.Context) {
	var reqBody models.UpdateTemplate

	if err := c.ShouldBindBodyWith(&reqBody, binding.JSON); err != nil {
		ResponseWithError(c, http.StatusBadRequest, ParseErrorMessage(err))
		return
	}

	templateID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid template ID")
		ResponseWithError(c, http.StatusBadRequest, "Template query error!")
	}

	template, err := h.DBService.GetTemplateByID(uint(templateID))
	if err != nil {
		log.Err(err).Msg("Template query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Template query error!")
		return
	}
	if template == nil {
		ResponseNotFound(c, "Template not found!")
		return
	}

	err = h.DBService.UpdateTemplate(uint(templateID), &reqBody)
	if err != nil {
		log.Err(err).Msg("Template update error")
		ResponseWithError(c, http.StatusInternalServerError, "Template updating error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Template updated!")
}

func (h *GinHandler) DeleteTemplate(c *gin.Context) {
	templateID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid template ID")
		ResponseWithError(c, http.StatusInternalServerError, "Invalid template ID")
		return
	}

	template, err := h.DBService.GetTemplateByID(uint(templateID))
	if err != nil {
		log.Err(err).Msg("Template query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Template query error!")
		return
	}
	if template == nil {
		ResponseNotFound(c, "Template not found!")
		return
	}

	err = h.DBService.DeleteTemplateByID(uint(templateID))
	if err != nil {
		log.Err(err).Msg("Template delete error")
		ResponseWithError(c, http.StatusInternalServerError, "Template delete error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Template deleted!")
}

func (h *GinHandler) GetTemplate(c *gin.Context) {
	templateID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Err(err)
	}
	template, err := h.DBService.GetTemplateByID(uint(templateID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Template query error!")
		return
	}
	if template == nil {
		ResponseNotFound(c, "Template not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Template details", template)
}

func (h *GinHandler) GetTemplates(c *gin.Context) {
	header := middleware.GetAuth(c)
	templates, err := h.DBService.GetTemplates(uint(header.UserID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Templates query error!")
		return
	}
	if templates == nil {
		ResponseNotFound(c, "Templates not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Template details", templates)
}
