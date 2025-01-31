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

func (s *SynMailServices) GetUserByEmailOrUserName(loginData *models.LoginRequest) (*models.User, error) {
	var user models.User

	where := map[string]interface{}{}
	if loginData.EmailID != "" {
		where["email_id"] = loginData.EmailID
	}

	if loginData.UserName != "" {
		where["user_name"] = loginData.UserName
	}

	if err := s.DB.Where(where).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
