package services

import "github.com/sahidhossen/synmail/src/models"

func (s *SynMailServices) CreateTemplate(template *models.Template) (*models.Template, error) {
	if err := s.DB.Create(&template).Error; err != nil {
		return nil, err
	}
	return template, nil
}

func (s *SynMailServices) GetTemplateByID(id uint) (*models.Template, error) {
	var template models.Template
	if err := s.DB.First(&template, id).Error; err != nil {
		return nil, err
	}
	return &template, nil
}

func (s *SynMailServices) GetTemplates(userID uint) ([]*models.Template, error) {
	var templates []*models.Template
	if err := s.DB.Where("user_id = ?", userID).First(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}

func (s *SynMailServices) DeleteTemplateByID(id uint) error {
	if err := s.DB.Delete(&models.Template{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (s *SynMailServices) UpdateTemplate(id uint, reqFields *models.UpdateTemplate) error {
	if err := s.DB.Model(&models.Template{}).Where("id = ?", id).Updates(reqFields).Error; err != nil {
		return err
	}
	return nil
}
