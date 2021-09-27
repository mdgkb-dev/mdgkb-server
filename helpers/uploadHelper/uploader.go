package uploadHelper

import (
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Uploader interface {
	GetUploaderPath() *string
	GetFullPath(*string) *string
	Upload(*gin.Context, []*multipart.FileHeader, *string) error
}

type LocalUploader struct {
	UploadPath *string
}

func NewLocalUploader(path *string) *LocalUploader {
	staticPath := filepath.Join(*path)
	return &LocalUploader{
		UploadPath: &staticPath,
	}
}

func (u *LocalUploader) Upload(c *gin.Context, file []*multipart.FileHeader, path *string) (err error) {
	fmt.Println(11)
	if path == nil {
		return errors.New("file does not relate to anything")
	}
	fmt.Println(22)
	uploadPath := u.GetUploaderPath()
	fmt.Println(*uploadPath)
	fullPath := filepath.Join(*uploadPath, *path)
	parts := strings.Split(fullPath, string(os.PathSeparator))
	fmt.Println(parts)
	err = os.MkdirAll("/"+filepath.Join(parts[:len(parts)-1]...), os.ModePerm)
	if err != nil {
		return err
	}

	err = c.SaveUploadedFile(file[0], fullPath)
	if err != nil {
		return err
	}
	return nil
}

func (u *LocalUploader) GetUploaderPath() *string {
	return u.UploadPath
}

func (u *LocalUploader) GetFullPath(path *string) *string {
	basePath := u.GetUploaderPath()
	fmt.Println(basePath, path)
	fullPath := filepath.Join(*basePath, *path)
	return &fullPath
}

func BuildPath(idFile *string) string {
	fullPath := filepath.Join(randomString(), randomString(), *idFile)
	return fullPath
}

func randomString() string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 1000
	return strconv.Itoa(rand.Intn(max-min+1) + min)
}
