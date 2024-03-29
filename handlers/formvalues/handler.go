package formvalues

import (
	"fmt"
	"net/http"

	"mdgkb/mdgkb-server/models"

	"github.com/pro-assistance/pro-assister/helpers/pdf"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.FormValue
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Upsert(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) UpdateMany(c *gin.Context) {
	var items models.FormValues
	err := c.Bind(&items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.UpsertMany(items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) DocumentsToPDF(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	files := item.GetFiles()
	filesToMerge := make(pdf.IFiles, 0)
	for i := range files {
		fullPath := h.helper.Uploader.GetFullPath(&files[i].FileSystemPath)
		files[i].FileSystemPath = *fullPath
		filesToMerge = append(filesToMerge, files[i])
	}
	mergedPDF, err := h.helper.PDF.MergeFilesToPDF(filesToMerge)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=response")
	c.Data(http.StatusOK, "application/pdf", mergedPDF)
}

func (h *Handler) DocumentsToZip(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	mergedPDF, err := h.filesService.FilesToZip(item.GetFiles())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.Header("Content-Description", "File Transfer")
	fileName := fmt.Sprintf("%s.zip", item.User.Human.GetFullName())
	h.helper.HTTP.SetFileHeaders(c, fileName)
	c.Data(http.StatusOK, "application/zip", mergedPDF)
}
