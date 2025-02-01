package handler

import (
	"encoding/csv"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/zerolog/log"
	"github.com/sahidhossen/synmail/src/middleware"
	"github.com/sahidhossen/synmail/src/models"
)

func (h *GinHandler) CreateSubscribe(c *gin.Context) {
	subscribe := &models.Subscriber{}

	if err := c.Bind(subscribe); err != nil {
		log.Err(err).Msg("Required field")
		ResponseWithError(c, http.StatusBadRequest, "Required field missing!")
		return
	}
	header := middleware.GetAuth(c)
	subscribe.UserID = header.UserID
	result, err := h.DBService.CreateSubscribe(subscribe)
	if err != nil {
		log.Err(err).Msg("Subscribe create error")
		ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, http.StatusOK, "success", "Subscribe created!", result)
}

func (h *GinHandler) ImportSubscriber(c *gin.Context) {
	// subscribe := &models.Subscriber{}

	file, err := c.FormFile("file")
	if err != nil {
		log.Err(err).Msg("Invalid file")
		ResponseWithError(c, http.StatusBadRequest, "Invalid file")
		return
	}

	// Validate file extension
	if ext := strings.ToLower(filepath.Ext(file.Filename)); ext != ".csv" {
		ResponseWithError(c, http.StatusBadRequest, "Only .csv files are allowed")
		return
	}

	// Open file for reading
	f, err := file.Open()
	if err != nil {
		log.Err(err).Msg("Failed to open file")
		ResponseWithError(c, http.StatusInternalServerError, "Failed to open file")
		return
	}
	defer f.Close()

	// Read CSV contents
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		log.Err(err).Msg("Failed to read CSV")
		ResponseWithError(c, http.StatusInternalServerError, "Failed to read CSV")
		return
	}

	// Convert CSV to JSON
	if len(records) < 1 {
		ResponseWithError(c, http.StatusBadRequest, "Empty CSV file")
		return
	}

	headers := records[0] // First row as headers
	var subscribers []*models.Subscriber
	auth := middleware.GetAuth(c)
	uniqueEmail := make(map[string]bool)
	for _, row := range records[1:] { // Skip headers
		entry := map[string]string{}
		for i, value := range row {
			entry[headers[i]] = value
		}
		_, pass := uniqueEmail[entry["Email"]]
		if !pass {
			uniqueEmail[entry["Email"]] = true
			subscribers = append(
				subscribers,
				&models.Subscriber{
					FirstName: entry["First Name"],
					LastName:  entry["Last Name"],
					Email:     entry["Email"],
					UserID:    auth.UserID,
				})
		}
	}
	err = h.DBService.CreateSubscribeInBatch(subscribers, 100)
	if err != nil {
		log.Err(err).Msg("Batch insert error")
		ResponseWithError(c, http.StatusInternalServerError, "Batch insert error")
		return
	}
	Response(c, http.StatusOK, "success", "CSV data inserted!", subscribers)
}

func (h *GinHandler) UpdateSubscribe(c *gin.Context) {
	var reqSubscribe models.UpdateSubscriber

	if err := c.ShouldBindBodyWith(&reqSubscribe, binding.JSON); err != nil {
		ResponseWithError(c, http.StatusBadRequest, ParseErrorMessage(err))
		return
	}

	subscribeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid subscribe ID")
		ResponseWithError(c, http.StatusBadRequest, "Subscribe query error!")
	}

	subscribe, err := h.DBService.GetSubscribeByID(uint(subscribeID))
	if err != nil {
		log.Err(err).Msg("Subscribe query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Subscribe query error!")
		return
	}
	if subscribe == nil {
		ResponseNotFound(c, "Subscribe not found!")
		return
	}

	err = h.DBService.UpdateSubscribe(uint(subscribeID), &reqSubscribe)
	if err != nil {
		log.Err(err).Msg("Subscribe update error")
		ResponseWithError(c, http.StatusInternalServerError, "Subscribe updating error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Subscribe updated!")
}

func (h *GinHandler) DeleteSubscribe(c *gin.Context) {
	subscribeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid subscribe ID")
		ResponseWithError(c, http.StatusInternalServerError, "Invalid subscribe ID")
		return
	}

	subscribe, err := h.DBService.GetSubscribeByID(uint(subscribeID))
	if err != nil {
		log.Err(err).Msg("Subscribe query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Subscribe query error!")
		return
	}
	if subscribe == nil {
		ResponseNotFound(c, "Subscribe not found!")
		return
	}

	err = h.DBService.DeleteSubscribeByID(uint(subscribeID))
	if err != nil {
		log.Err(err).Msg("Subscribe delete error")
		ResponseWithError(c, http.StatusInternalServerError, "Subscribe delete error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Subscribe deleted!")
}

func (h *GinHandler) GetSubscribe(c *gin.Context) {
	subscribeID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Err(err)
	}
	subscribe, err := h.DBService.GetSubscribeByID(uint(subscribeID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Subscribe query error!")
		return
	}
	if subscribe == nil {
		ResponseNotFound(c, "Subscribe not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Subscribe details", subscribe)
}

func (h *GinHandler) GetSubscribers(c *gin.Context) {
	header := middleware.GetAuth(c)

	subscribers, err := h.DBService.GetSubscribers(uint(header.UserID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Subscribers query error!")
		return
	}
	if subscribers == nil {
		ResponseNotFound(c, "Subscribers not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Subscriber list", subscribers)
}
