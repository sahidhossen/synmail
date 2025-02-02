package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
	"github.com/sahidhossen/synmail/src/models"
	"gorm.io/gorm"
)

func (h *GinHandler) CreateTopicMap(c *gin.Context) {
	topicMap := &models.SubscribeTopicMap{}

	if err := c.Bind(topicMap); err != nil {
		log.Err(err).Msg("Required field")
		ResponseWithError(c, http.StatusBadRequest, "Required field missing!")
		return
	}
	existingMap, err := h.DBService.GetSubscribeTopicMapByTopic(topicMap.TopicID, topicMap.SubscribeID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Err(err).Msg("Topic map query error")
			ResponseWithError(c, http.StatusBadRequest, "Topic map query error!")
			return
		}
	}
	if existingMap != nil {
		log.Err(err).Msg("Duplicate entry")
		ResponseWithError(c, http.StatusBadRequest, "Duplicate entry!")
		return
	}

	result, err := h.DBService.CreateTopicMap(topicMap)
	if err != nil {
		log.Err(err).Msg("Topic Map create error")
		ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, http.StatusOK, "success", "Topic Map created!", result)
}

func (h *GinHandler) UpdateTopicMap(c *gin.Context) {
	type TopicMapReqest struct {
		TopicID     uint `json:"topic_id"`
		SubscribeID uint `json:"subscribe_id"`
	}
	var reqTopic TopicMapReqest

	if err := c.ShouldBindBodyWith(&reqTopic, binding.JSON); err != nil {
		ResponseWithError(c, http.StatusBadRequest, ParseErrorMessage(err))
		return
	}

	topicMapID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid topic map ID")
		ResponseWithError(c, http.StatusBadRequest, "topic map query error!")
	}

	topic, err := h.DBService.GetSubscribeTopicMapByID(uint(topicMapID))
	if err != nil {
		log.Err(err).Msg("Topic query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Topic map query error!")
		return
	}
	if topic == nil {
		ResponseNotFound(c, "Topic map not found!")
		return
	}

	err = h.DBService.UpdateSubscribeTopicMap(
		uint(topicMapID),
		map[string]interface{}{"topic_id": reqTopic.TopicID, "subscribe_id": reqTopic.SubscribeID},
	)
	if err != nil {
		log.Err(err).Msg("Topic Map update error")
		ResponseWithError(c, http.StatusInternalServerError, "Topic Map updating error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Topic map updated!")
}

func (h *GinHandler) DeleteTopicMap(c *gin.Context) {
	topicID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid topic map ID")
		ResponseWithError(c, http.StatusInternalServerError, "Invalid topic map ID")
		return
	}

	topicMap, err := h.DBService.GetSubscribeTopicMapByID(uint(topicID))
	if err != nil {
		log.Err(err).Msg("Topic map query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Topic map query error!")
		return
	}
	if topicMap == nil {
		ResponseNotFound(c, "Topic map not found!")
		return
	}

	err = h.DBService.DeleteSubscribeTopicMapByID(uint(topicID))
	if err != nil {
		log.Err(err).Msg("Topic map delete error")
		ResponseWithError(c, http.StatusInternalServerError, "Topic map delete error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Topic map deleted!")
}

func (h *GinHandler) GetSubscribeTopicMap(c *gin.Context) {
	topicID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Err(err)
	}
	topic, err := h.DBService.GetSubscribeTopicMapByID(uint(topicID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Topic map query error!")
		return
	}
	if topic == nil {
		ResponseNotFound(c, "Topic not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Topic details", topic)
}
