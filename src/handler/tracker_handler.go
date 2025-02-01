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

func (h *GinHandler) CreateTracker(c *gin.Context) {
	tracker := &models.Trackers{}

	if err := c.Bind(tracker); err != nil {
		log.Err(err).Msg("Required field")
		ResponseWithError(c, http.StatusBadRequest, "Required field missing!")
		return
	}

	result, err := h.DBService.CreateTracker(tracker)
	if err != nil {
		log.Err(err).Msg("Tracker create error")
		ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, http.StatusOK, "success", "Tracker created!", result)
}

func (h *GinHandler) UpdateTracker(c *gin.Context) {

	var reqBody models.TrackersUpdate

	if err := c.ShouldBindBodyWith(&reqBody, binding.JSON); err != nil {
		ResponseWithError(c, http.StatusBadRequest, ParseErrorMessage(err))
		return
	}

	trackerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid tracker ID")
		ResponseWithError(c, http.StatusBadRequest, "Tracker query error!")
	}

	tracker, err := h.DBService.GetTrackerByID(uint(trackerID))
	if err != nil {
		log.Err(err).Msg("Tracker query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Tracker query error!")
		return
	}
	if tracker == nil {
		ResponseNotFound(c, "Tracker not found!")
		return
	}

	err = h.DBService.UpdateTracker(uint(trackerID), &reqBody)
	if err != nil {
		log.Err(err).Msg("Tracker update error")
		ResponseWithError(c, http.StatusInternalServerError, "Tracker updating error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Tracker updated!")
}

func (h *GinHandler) DeleteTracker(c *gin.Context) {
	trackerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid tracker ID")
		ResponseWithError(c, http.StatusInternalServerError, "Invalid tracker ID")
		return
	}

	tracker, err := h.DBService.GetTrackerByID(uint(trackerID))
	if err != nil {
		log.Err(err).Msg("Tracker query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Tracker query error!")
		return
	}
	if tracker == nil {
		ResponseNotFound(c, "Tracker not found!")
		return
	}

	err = h.DBService.DeleteTrackerByID(uint(trackerID))
	if err != nil {
		log.Err(err).Msg("Tracker delete error")
		ResponseWithError(c, http.StatusInternalServerError, "Tracker delete error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Tracker deleted!")
}

func (h *GinHandler) GetTracker(c *gin.Context) {
	trackerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Err(err)
	}
	tracker, err := h.DBService.GetTrackerByID(uint(trackerID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Tracker query error!")
		return
	}
	if tracker == nil {
		ResponseNotFound(c, "Tracker not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Tracker details", tracker)
}

func (h *GinHandler) GetTrackers(c *gin.Context) {
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
	Response(c, http.StatusOK, "success", "Tracker details", trackers)
}
