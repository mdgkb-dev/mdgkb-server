package helpers

import (
	"mdgkb/mdgkb-server/config"
	httpHelper "mdgkb/mdgkb-server/helpers/httpHelperV2"
	"mdgkb/mdgkb-server/helpers/pdfHelper"
	"mdgkb/mdgkb-server/helpers/sqlHelper"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

type Helper struct {
	HTTP *httpHelper.HTTPHelper
	PDF *pdfHelper.PDFHelper
	Uploader uploadHelper.Uploader
	SQL *sqlHelper.SQLHelper
}

func NewHelper(config config.Config) *Helper {
	http := httpHelper.NewHTTPHelper()
	pdf := pdfHelper.NewPDFHelper(config)
	sql := sqlHelper.NewSQLHelper()
	return &Helper{HTTP: http, Uploader: uploadHelper.NewLocalUploader(&config.UploadPath), PDF: pdf, SQL: sql}
}