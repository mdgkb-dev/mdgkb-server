package dpoApplications

import (
	"context"
	"github.com/pro-assistance/pro-assister/helper"
	"github.com/pro-assistance/pro-assister/sqlHelper"
	"mdgkb/mdgkb-server/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IHandler interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	EmailExists(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)

	SubscribeCreate(c *gin.Context)
}

type IService interface {
	setQueryFilter(*gin.Context) error
	GetAll() (models.DpoApplications, error)
	Get(*string) (*models.DpoApplication, error)
	EmailExists(string, string) (bool, error)
	Create(*models.DpoApplication) error
	Update(*models.DpoApplication) error
	Delete(*string) error
}

type IRepository interface {
	setQueryFilter(*gin.Context) error
	getDB() *bun.DB
	getAll() (models.DpoApplications, error)
	get(*string) (*models.DpoApplication, error)
	emailExists(string, string) (bool, error)
	create(*models.DpoApplication) error
	update(*models.DpoApplication) error
	delete(*string) error
}

type IFilesService interface {
	Upload(*gin.Context, *models.DpoApplication, map[string][]*multipart.FileHeader) error
}

type Event struct {
	Message       chan string
	NewClients    chan chan string
	ClosedClients chan chan string
	TotalClients  map[chan string]bool
}

func (stream *Event) listen() {
	for {
		select {
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

type Handler struct {
	service      IService
	filesService IFilesService
	helper       *helper.Helper
	sse          *Event
}

type Service struct {
	repository IRepository
	helper     *helper.Helper
}

type Repository struct {
	db          *bun.DB
	ctx         context.Context
	helper      *helper.Helper
	queryFilter *sqlHelper.QueryFilter
}

type FilesService struct {
	helper *helper.Helper
}

func CreateHandler(db *bun.DB, helper *helper.Helper) *Handler {
	repo := NewRepository(db, helper)
	service := NewService(repo, helper)
	filesService := NewFilesService(helper)
	return NewHandler(service, filesService, helper)
}

// NewHandler constructor
func NewHandler(s IService, filesService IFilesService, helper *helper.Helper) *Handler {
	event := NewServer()
	return &Handler{service: s, filesService: filesService, helper: helper, sse: event}
}

func NewService(repository IRepository, helper *helper.Helper) *Service {
	return &Service{repository: repository, helper: helper}
}

func NewRepository(db *bun.DB, helper *helper.Helper) *Repository {
	return &Repository{db: db, ctx: context.Background(), helper: helper}
}

func NewFilesService(helper *helper.Helper) *FilesService {
	return &FilesService{helper: helper}
}

func NewServer() (event *Event) {
	event = &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}
	go event.listen()
	return
}

// New event messages are broadcast to all registered client connection channels
type ClientChan chan string

func (stream *Event) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientChan := make(ClientChan)
		stream.NewClients <- clientChan
		defer func() {
			stream.ClosedClients <- clientChan
		}()
		go func() {
			<-c.Done()
			stream.ClosedClients <- clientChan
		}()
		c.Next()
	}
}
