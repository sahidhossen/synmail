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

func (h *GinHandler) CreateTopics(c *gin.Context) {
	topic := &models.SubscribeTopics{}

	if err := c.Bind(topic); err != nil {
		log.Err(err).Msg("Required field")
		ResponseWithError(c, http.StatusBadRequest, "Required field missing!")
		return
	}

	result, err := h.DBService.CreateTopic(topic)
	if err != nil {
		log.Err(err).Msg("Topic create error")
		ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, http.StatusOK, "success", "Topic created!", result)
}

func (h *GinHandler) UpdateTopic(c *gin.Context) {
	type TopicReqest struct {
		Name string `json:"name" binding:"required"`
	}
	var reqTopic TopicReqest

	if err := c.ShouldBindBodyWith(&reqTopic, binding.JSON); err != nil {
		ResponseWithError(c, http.StatusBadRequest, ParseErrorMessage(err))
		return
	}

	topicID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid topic ID")
		ResponseWithError(c, http.StatusBadRequest, "topic query error!")
	}

	topic, err := h.DBService.GetSubscribeTopicByID(uint(topicID))
	if err != nil {
		log.Err(err).Msg("Topic query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Topic query error!")
		return
	}
	if topic == nil {
		ResponseNotFound(c, "Topic not found!")
		return
	}

	err = h.DBService.UpdateSubscribeTopic(uint(topicID), reqTopic.Name)
	if err != nil {
		log.Err(err).Msg("Topic update error")
		ResponseWithError(c, http.StatusInternalServerError, "Topic updating error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Topic updated!")
}

func (h *GinHandler) DeleteTopic(c *gin.Context) {
	topicID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid topic ID")
		ResponseWithError(c, http.StatusInternalServerError, "Invalid topic ID")
		return
	}

	topic, err := h.DBService.GetSubscribeTopicByID(uint(topicID))
	if err != nil {
		log.Err(err).Msg("Topic query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Topic query error!")
		return
	}
	if topic == nil {
		ResponseNotFound(c, "Topic not found!")
		return
	}

	err = h.DBService.DeleteSubscribeTopicByID(uint(topicID))
	if err != nil {
		log.Err(err).Msg("Topic delete error")
		ResponseWithError(c, http.StatusInternalServerError, "Topic delete error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Topic deleted!")
}

func (h *GinHandler) GetSubscribeTopic(c *gin.Context) {
	topicID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Err(err)
	}
	topic, err := h.DBService.GetSubscribeTopicByID(uint(topicID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Topic query error!")
		return
	}
	if topic == nil {
		ResponseNotFound(c, "Topic not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Topic details", topic)
}

func (h *GinHandler) GetSubscribeTopices(c *gin.Context) {
	header := middleware.GetAuth(c)

	trackers, err := h.DBService.GetTrackers(uint(header.UserID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Trackers query error!")
		return
	}
	if trackers == nil {
		ResponseNotFound(c, "Trackers not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Subscriber list", trackers)
}
