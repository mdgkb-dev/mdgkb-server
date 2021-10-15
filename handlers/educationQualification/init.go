package educationQualification

type IHandler interface {
GetAll(c *gin.Context) error
Get(c *gin.Context) error
Create(c *gin.Context) error
Update(c *gin.Context) error
Delete(c *gin.Context) error
}

type IService interface {
GetAll() (models.EducationQualifications, error)
Get(*string) (*models.EducationQualification, error)
Create(*models.EducationQualification) error
Update(*models.EducationQualification) error
Delete(*string) error
}

type IRepository interface {
create(*models.EducationQualification) error
getAll() (models.EducationQualifications, error)
get(*string) (*models.EducationQualification, error)
update(*models.EducationQualification) error
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
