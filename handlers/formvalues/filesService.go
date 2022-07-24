package formvalues

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(c *gin.Context, item *models.FormValue, files map[string][]*multipart.FileHeader) (err error) {
	for i, file := range files {
		err = s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *FilesService) FilesToZip(fileInfos models.FileInfos) ([]byte, error) {
	return s.WriteZipBytes(fileInfos.GetPathsAndNames())
}

// WriteZipBytes func
func (s *FilesService) WriteZipBytes(paths, names []string) ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)

	for i, file := range paths {
		data, err := ioutil.ReadFile(filepath.Join(*s.helper.Uploader.GetUploaderPath(), file))
		if err != nil {
			return nil, err
		}
		f, err := writer.Create(filepath.Base(names[i]))
		if err != nil {
			return nil, err
		}
		_, err = f.Write(data)
		if err != nil {
			return nil, err
		}
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
