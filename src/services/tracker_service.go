package services

import "github.com/sahidhossen/synmail/src/models"

func (s *SynMailServices) CreateTracker(tracker *models.Trackers) (*models.Trackers, error) {
	if err := s.DB.Create(&tracker).Error; err != nil {
		return nil, err
	}
	return tracker, nil
}

func (s *SynMailServices) GetTrackerByID(id uint) (*models.Trackers, error) {
	var tracker models.Trackers
	if err := s.DB.First(&tracker, id).Error; err != nil {
		return nil, err
	}
	return &tracker, nil
}

func (s *SynMailServices) GetTrackers(userID uint) ([]*models.Trackers, error) {
	var trackers []*models.Trackers
	if err := s.DB.Where("user_id = ?", userID).First(&trackers).Error; err != nil {
		return nil, err
	}
	return trackers, nil
}

func (s *SynMailServices) DeleteTrackerByID(id uint) error {
	if err := s.DB.Delete(&models.Trackers{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SynMailServices) UpdateTracker(id uint, reqFields *models.TrackersUpdate) error {
	if err := s.DB.Model(&models.Trackers{}).Where("id = ?", id).Updates(reqFields).Error; err != nil {
		return err
	}
	return nil
}
