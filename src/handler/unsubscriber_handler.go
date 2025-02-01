package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
	"github.com/sahidhossen/synmail/src/models"
)

func (h *GinHandler) CreateUnsubscribe(c *gin.Context) {
	subscribe := &models.Unsubscribers{}

	if err := c.Bind(subscribe); err != nil {
		log.Err(err).Msg("Required field")
		ResponseWithError(c, http.StatusBadRequest, "Required field missing!")
		return
	}

	result, err := h.DBService.CreateUnSubscriber(subscribe)
	if err != nil {
		log.Err(err).Msg("Unsubscribe create error")
		ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, http.StatusOK, "success", "Unsubscribe created!", result)
}

func (h *GinHandler) UpdateUnsubscribe(c *gin.Context) {

	var reqBody models.UnsubscribeUpdate

	if err := c.ShouldBindBodyWith(&reqBody, binding.JSON); err != nil {
		ResponseWithError(c, http.StatusBadRequest, ParseErrorMessage(err))
		return
	}

	subscribeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid subscribe ID")
		ResponseWithError(c, http.StatusBadRequest, "Subscribe query error!")
	}

	subscribe, err := h.DBService.GetUnSubscribeByID(uint(subscribeID))
	if err != nil {
		log.Err(err).Msg("Unsubscribe query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Unsubscribe query error!")
		return
	}
	if subscribe == nil {
		ResponseNotFound(c, "Unsubscribe not found!")
		return
	}

	err = h.DBService.UpdateUnSubscribe(uint(subscribeID), &reqBody)
	if err != nil {
		log.Err(err).Msg("Unsubscribe update error")
		ResponseWithError(c, http.StatusInternalServerError, "Unsubscribe updating error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Unsubscribe updated!")
}

func (h *GinHandler) DeleteUnsubscribe(c *gin.Context) {
	subscribeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid Unsubscribe ID")
		ResponseWithError(c, http.StatusInternalServerError, "Invalid Unsubscribe ID")
		return
	}

	subscribe, err := h.DBService.GetUnSubscribeByID(uint(subscribeID))
	if err != nil {
		log.Err(err).Msg("Unsubscribe query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Unsubscribe query error!")
		return
	}
	if subscribe == nil {
		ResponseNotFound(c, "Unsubscribe not found!")
		return
	}

	err = h.DBService.DeleteUnSubscribeByID(uint(subscribeID))
	if err != nil {
		log.Err(err).Msg("Unsubscribe delete error")
		ResponseWithError(c, http.StatusInternalServerError, "Unsubscribe delete error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Unsubscribe deleted!")
}

func (h *GinHandler) GetUnSubscribe(c *gin.Context) {
	subscribeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Err(err)
	}
	subscribe, err := h.DBService.GetUnSubscribeByID(uint(subscribeID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Unsubscribe query error!")
		return
	}
	if subscribe == nil {
		ResponseNotFound(c, "Unsubscribe not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Unsubscribe details", subscribe)
}
