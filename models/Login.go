package models

type Login struct {
	Email    string `json:"email" validate:"required,omitempty"`
	Password string `json:"password" validate:"required,omitempty"`
}
