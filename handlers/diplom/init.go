package diplom

type IHandler interface {
GetAll(c *gin.Context) error
Get(c *gin.Context) error
Create(c *gin.Context) error
Update(c *gin.Context) error
Delete(c *gin.Context) error
}

type IService interface {
GetAll() (models.Diploms, error)
Get(*string) (*models.Diplom, error)
Create(*models.Diplom) error
Update(*models.Diplom) error
Delete(*string) error
}

type IRepository interface {
create(*models.Diplom) error
getAll() (models.Diploms, error)
get(*string) (*models.Diplom, error)
update(*models.Diplom) error
delete(*string) error
}

type Handler struct {
service IService
}

type Service struct {
repository IRepository
}

type Repository struct {
db  *bun.DB
ctx context.Context
}

func CreateHandler(db *bun.DB) *Handler {
repo := NewRepository(db)
service := NewService(repo)
return NewHandler(service)
}

// NewHandler constructor
func NewHandler(s IService) *Handler {
return &Handler{service: s}
}

func NewService(repository IRepository) *Service {
return &Service{repository: repository}
}

func NewRepository(db *bun.DB) *Repository {
return &Repository{db: db, ctx: context.Background()}
}
