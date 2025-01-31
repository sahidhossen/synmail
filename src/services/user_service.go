package services

import "github.com/sahidhossen/synmail/src/models"

func (s *SynMailServices) CreateUser(user *models.User) error {
	if err := s.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *SynMailServices) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
