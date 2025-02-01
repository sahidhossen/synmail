package services

import "github.com/sahidhossen/synmail/src/models"

func (s *SynMailServices) CreateSubscribe(subscribe *models.Subscriber) (*models.Subscriber, error) {
	if err := s.DB.Create(&subscribe).Error; err != nil {
		return nil, err
	}
	return subscribe, nil
}

func (s *SynMailServices) CreateSubscribeInBatch(subscribers []*models.Subscriber, batchSize int) error {
	if err := s.DB.CreateInBatches(&subscribers, batchSize).Error; err != nil {
		return err
	}
	return nil
}

func (s *SynMailServices) GetSubscribeByID(id uint) (*models.Subscriber, error) {
	var subscribe models.Subscriber
	if err := s.DB.First(&subscribe, id).Error; err != nil {
		return nil, err
	}
	return &subscribe, nil
}

func (s *SynMailServices) GetSubscribers(userID uint) ([]*models.Subscriber, error) {
	var subscribers []*models.Subscriber
	if err := s.DB.Where("user_id = ?", userID).Find(&subscribers).Error; err != nil {
		return nil, err
	}
	return subscribers, nil
}

func (s *SynMailServices) DeleteSubscribeByID(id uint) error {
	if err := s.DB.Delete(&models.Subscriber{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SynMailServices) UpdateSubscribe(id uint, reqFields *models.UpdateSubscriber) error {
	if err := s.DB.Model(&models.Subscriber{}).Where("id = ?", id).Updates(reqFields).Error; err != nil {
		return err
	}
	return nil
}
