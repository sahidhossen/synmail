package services

import "github.com/sahidhossen/synmail/src/models"

func (s *SynMailServices) CreateTopic(subscribeTopic *models.SubscribeTopics) (*models.SubscribeTopics, error) {
	if err := s.DB.Create(&subscribeTopic).Error; err != nil {
		return nil, err
	}
	return subscribeTopic, nil
}

func (s *SynMailServices) GetSubscribeTopicByID(id uint) (*models.SubscribeTopics, error) {
	var subscribeTopic models.SubscribeTopics
	if err := s.DB.First(&subscribeTopic, id).Error; err != nil {
		return nil, err
	}
	return &subscribeTopic, nil
}

func (s *SynMailServices) DeleteSubscribeTopicByID(id uint) error {
	if err := s.DB.Delete(&models.SubscribeTopics{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SynMailServices) UpdateSubscribeTopic(id uint, name string) error {
	if err := s.DB.Model(&models.SubscribeTopics{}).Where("id = ?", id).Updates(map[string]interface{}{"name": name}).Error; err != nil {
		return err
	}
	return nil
}
