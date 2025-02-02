package services

import "github.com/sahidhossen/synmail/src/models"

func (s *SynMailServices) CreateTopicMap(subscribeTopicMap *models.SubscribeTopicMap) (*models.SubscribeTopicMap, error) {
	if err := s.DB.Create(&subscribeTopicMap).Error; err != nil {
		return nil, err
	}
	return subscribeTopicMap, nil
}

func (s *SynMailServices) GetSubscribeTopicMapByID(id uint) (*models.SubscribeTopicMap, error) {
	var subscribeTopicMap models.SubscribeTopicMap
	if err := s.DB.First(&subscribeTopicMap, id).Error; err != nil {
		return nil, err
	}
	return &subscribeTopicMap, nil
}

func (s *SynMailServices) GetSubscribeTopicMapByTopic(topicID, subscriberID uint) (*models.SubscribeTopicMap, error) {
	var subscribeTopicMap models.SubscribeTopicMap
	if err := s.DB.Where("topic_id = ? AND subscribe_id = ?", topicID, subscriberID).First(&subscribeTopicMap).Error; err != nil {
		return nil, err
	}
	return &subscribeTopicMap, nil
}

func (s *SynMailServices) DeleteSubscribeTopicMapByID(id uint) error {
	if err := s.DB.Delete(&models.SubscribeTopicMap{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SynMailServices) UpdateSubscribeTopicMap(id uint, fields map[string]interface{}) error {
	if err := s.DB.Model(&models.SubscribeTopicMap{}).Where("id = ?", id).Updates(fields).Error; err != nil {
		return err
	}
	return nil
}
