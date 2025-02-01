package services

import "github.com/sahidhossen/synmail/src/models"

func (s *SynMailServices) CreateUnSubscriber(unsubscribe *models.Unsubscribers) (*models.Unsubscribers, error) {
	if err := s.DB.Create(&unsubscribe).Error; err != nil {
		return nil, err
	}
	return unsubscribe, nil
}

func (s *SynMailServices) GetUnSubscribeByID(id uint) (*models.Unsubscribers, error) {
	var unsubscribe models.Unsubscribers
	if err := s.DB.First(&unsubscribe, id).Error; err != nil {
		return nil, err
	}
	return &unsubscribe, nil
}

func (s *SynMailServices) GetUnsubscribers(userID uint) ([]*models.Unsubscribers, error) {
	var unsubscribe []*models.Unsubscribers
	if err := s.DB.Where("user_id = ?", userID).Find(&unsubscribe).Error; err != nil {
		return nil, err
	}
	return unsubscribe, nil
}

func (s *SynMailServices) DeleteUnSubscribeByID(id uint) error {
	if err := s.DB.Delete(&models.Unsubscribers{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SynMailServices) UpdateUnSubscribe(id uint, reqFields *models.UnsubscribeUpdate) error {
	if err := s.DB.Model(&models.Trackers{}).Where("id = ?", id).Updates(reqFields).Error; err != nil {
		return err
	}
	return nil
}
