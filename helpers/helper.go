package helpers

import (
	configPack "mdgkb/mdgkb-server/config"
	"mdgkb/mdgkb-server/helpers/elasticSearchHelper"
	"mdgkb/mdgkb-server/helpers/emailHelper"
	httpHelper "mdgkb/mdgkb-server/helpers/httpHelperV2"
	"mdgkb/mdgkb-server/helpers/pdfHelper"
	"mdgkb/mdgkb-server/helpers/socialHelper"
	"mdgkb/mdgkb-server/helpers/sqlHelper"
	"mdgkb/mdgkb-server/helpers/tokenHelper"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

type Helper struct {
	HTTP     *httpHelper.HTTPHelper
	Search   *elasticSearchHelper.ElasticSearchHelper
	PDF      *pdfHelper.PDFHelper
	Uploader uploadHelper.Uploader
	SQL      *sqlHelper.SQLHelper
	Token    *tokenHelper.TokenHelper
	Email    *emailHelper.EmailHelper
	Social   *socialHelper.Social
}

func NewHelper(config configPack.Config) *Helper {
	http := httpHelper.NewHTTPHelper()
	pdf := pdfHelper.NewPDFHelper(config)
	sql := sqlHelper.NewSQLHelper()
	uploader := uploadHelper.NewLocalUploader(&config.UploadPath)
	token := tokenHelper.NewTokenHelper(config.TokenSecret)
	email := emailHelper.NewEmailHelper(config.Email)
	social := socialHelper.NewSocial(config.Social)
	search := elasticSearchHelper.NewElasticSearchHelper(config.ElasticSearch.ElasticSearchOn)
	return &Helper{HTTP: http, Uploader: uploader, PDF: pdf, SQL: sql, Token: token, Email: email, Social: social, Search: search}
}
