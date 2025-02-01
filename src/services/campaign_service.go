package services

import "github.com/sahidhossen/synmail/src/models"

func (s *SynMailServices) CreateCampaign(campaign *models.Campaign) (*models.Campaign, error) {
	if err := s.DB.Create(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s *SynMailServices) GetCampaignByID(id uint) (*models.Campaign, error) {
	var campaign models.Campaign
	if err := s.DB.First(&campaign, id).Error; err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (s *SynMailServices) GetCampaignByUserID(userID uint) ([]models.Campaign, error) {
	var campaign []models.Campaign
	if err := s.DB.Where("user_id = ?", userID).Find(&campaign).Error; err != nil {
		return nil, err
	}
	return campaign, nil
}

func (s *SynMailServices) DeleteCampaignByID(id uint) error {
	if err := s.DB.Delete(&models.Campaign{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SynMailServices) UpdateCampaign(id uint, reqCampaign *models.UpdateCampaign) error {
	if err := s.DB.Model(&models.Campaign{}).Where("id = ?", id).Updates(reqCampaign).Error; err != nil {
		return err
	}
	return nil
}
