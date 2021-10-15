package educationType

type IHandler interface {
GetAll(c *gin.Context) error
Get(c *gin.Context) error
Create(c *gin.Context) error
Update(c *gin.Context) error
Delete(c *gin.Context) error
}

type IService interface {
GetAll() (models.EducationTypes, error)
Get(*string) (*models.EducationType, error)
Create(*models.EducationType) error
Update(*models.EducationType) error
Delete(*string) error
}

type IRepository interface {
create(*models.EducationType) error
getAll() (models.EducationTypes, error)
get(*string) (*models.EducationType, error)
update(*models.EducationType) error
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
