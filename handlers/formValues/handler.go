package formValues

import (
	"fmt"
	"github.com/pro-assistance/pro-assister/pdfHelper"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(&id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.FormValue
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Upsert(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) DocumentsToPDF(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(&id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	files := item.GetFiles()
	filesToMerge := make(pdfHelper.IFiles, 0)
	for i := range files {
		fullPath := h.helper.Uploader.GetFullPath(&files[i].FileSystemPath)
		files[i].FileSystemPath = *fullPath
		fmt.Println(files[i].GetFullPath())
		filesToMerge = append(filesToMerge, files[i])

	}
	mergedPDF, err := h.helper.PDF.MergeFilesToPDF(filesToMerge)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=response")
	c.Data(http.StatusOK, "application/pdf", mergedPDF)
}
