package helpers

import (
	"mdgkb/mdgkb-server/config"
	httpHelper "mdgkb/mdgkb-server/helpers/httpHelperV2"
	"mdgkb/mdgkb-server/helpers/pdfHelper"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

type Helper struct {
	HTTP *httpHelper.HTTPHelper
	PDF *pdfHelper.PDFHelper
	Uploader uploadHelper.Uploader
}

func NewHelper(config config.Config) *Helper {
	http := httpHelper.NewHTTPHelper()
	pdf := pdfHelper.NewPDFHelper(config)
	return &Helper{HTTP: http, Uploader: uploadHelper.NewLocalUploader(&config.UploadPath), PDF: pdf}
}