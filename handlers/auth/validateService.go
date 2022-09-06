package auth

import (
	"mdgkb/mdgkb-server/models"
)

func (s *ValidateService) Login(item *models.Login) error {
	return s.helper.Validator.Validate(item)
}
