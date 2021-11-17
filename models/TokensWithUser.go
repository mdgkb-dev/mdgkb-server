package models

import "mdgkb/mdgkb-server/helpers"

type TokensWithUser struct {
	Token *helpers.TokenDetails `json:"token"`
	User  User           `json:"user"`
}
