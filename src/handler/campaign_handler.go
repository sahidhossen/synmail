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

func (h *GinHandler) CreateCampaign(c *gin.Context) {
	campaign := &models.Campaign{}

	if err := c.Bind(campaign); err != nil {
		log.Err(err).Msg("Required field")
		ResponseWithError(c, http.StatusBadRequest, "Required field missing!")
		return
	}

	header := middleware.GetAuth(c)
	campaign.UserID = header.UserID

	result, err := h.DBService.CreateCampaign(campaign)
	if err != nil {
		log.Err(err).Msg("campaign create error")
		ResponseWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	Response(c, http.StatusOK, "success", "Campaign created!", result)
}

func (h *GinHandler) UpdateCampaign(c *gin.Context) {
	var reqCampaign models.UpdateCampaign

	if err := c.ShouldBindBodyWith(&reqCampaign, binding.JSON); err != nil {
		ResponseWithError(c, http.StatusBadRequest, ParseErrorMessage(err))
		return
	}

	if reqCampaign == (models.UpdateCampaign{}) {
		ResponseWithError(c, http.StatusBadRequest, "Request body empty!")
		return
	}

	campaignID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid campaign ID")
		ResponseWithError(c, http.StatusBadRequest, "Campaign query error!")
	}

	campaign, err := h.DBService.GetCampaignByID(uint(campaignID))
	if err != nil {
		log.Err(err).Msg("Campaign query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Campaign query error!")
		return
	}
	if campaign == nil {
		ResponseNotFound(c, "Campaign not found!")
		return
	}

	err = h.DBService.UpdateCampaign(uint(campaignID), &reqCampaign)
	if err != nil {
		log.Err(err).Msg("Campaign update error")
		ResponseWithError(c, http.StatusInternalServerError, "Campaign updating error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Campaign updated!")
}

func (h *GinHandler) DeleteCampaign(c *gin.Context) {
	campaignID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Err(err).Msg("Invalid campaign ID")
		ResponseWithError(c, http.StatusInternalServerError, "Invalid campaign ID")
		return
	}

	campaign, err := h.DBService.GetCampaignByID(uint(campaignID))
	if err != nil {
		log.Err(err).Msg("Campaign query error!")
		ResponseWithError(c, http.StatusInternalServerError, "Campaign query error!")
		return
	}
	if campaign == nil {
		ResponseNotFound(c, "Campaign not found!")
		return
	}

	err = h.DBService.DeleteCampaignByID(uint(campaignID))
	if err != nil {
		log.Err(err).Msg("Campaign delete error")
		ResponseWithError(c, http.StatusInternalServerError, "Campaign delete error!")
		return
	}
	ResponseWithMsg(c, http.StatusOK, "success", "Campaign deleted!")
}

func (h *GinHandler) GetCampaign(c *gin.Context) {
	campaignID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Err(err)
	}
	campaign, err := h.DBService.GetCampaignByID(uint(campaignID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Campaign query error!")
		return
	}
	if campaign == nil {
		ResponseNotFound(c, "Campaign not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Campaign details", campaign)
}

func (h *GinHandler) GetCampaigns(c *gin.Context) {
	header := middleware.GetAuth(c)
	campaigns, err := h.DBService.GetCampaignByUserID(uint(header.UserID))
	if err != nil {
		log.Err(err)
		ResponseWithError(c, http.StatusInternalServerError, "Campaigns query error!")
		return
	}
	if campaigns == nil {
		ResponseNotFound(c, "Campaigns not found!")
		return
	}
	Response(c, http.StatusOK, "success", "Campaign list", campaigns)
}
