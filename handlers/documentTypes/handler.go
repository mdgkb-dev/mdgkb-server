package documentTypes

import (
	"mdgkb/mdgkb-server/helpers/httpHelper"
	"mdgkb/mdgkb-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.DocumentType
	_, err := httpHelper.GetForm(c, &item)
	//if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//err = h.filesService.Upload(c, &item, files)
	err = h.service.Create(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll(models.CreateDocumentsParamsFromContext(c))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(&id)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(&id)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.DocumentType
	_, err := httpHelper.GetForm(c, &item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	//err = h.filesService.Upload(c, &item, files)

	err = h.service.Update(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetDocumentsTypesForTablesNames(c *gin.Context) {
	items := h.service.GetDocumentsTypesForTablesNames()
	c.JSON(http.StatusOK, items)
}