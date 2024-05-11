package auth

import (
	"github.com/pro-assistance/pro-assister/helper"
)

type Handler struct {
	helper *helper.Helper
}

type Service struct {
	helper *helper.Helper
}

type Repository struct {
	helper *helper.Helper
}

type FilesService struct {
	helper *helper.Helper
}

type ValidateService struct {
	helper *helper.Helper
}

type DoesLoginExist struct {
	DoesLoginExist bool
}

var (
	H *Handler
	S *Service
)

func Init(h *helper.Helper) {
	H = &Handler{helper: h}
	S = &Service{helper: h}
}
