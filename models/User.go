package models

import (
	"time"

	"github.com/pro-assistance/pro-assister/middleware"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"users,select:users_view,alias:users_view"`

	ID                uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Email             string        `json:"email"`
	UUID              uuid.UUID     `bun:"type:uuid,nullzero,notnull,default:uuid_generate_v4()"  json:"uuid"` // для восстановления пароля - обеспечивает уникальность страницы на фронте
	Phone             string        `json:"phone"`
	Password          string        `json:"password"`
	IsActive          bool          `json:"isActive"`
	Human             *Human        `bun:"rel:belongs-to" json:"human"`
	HumanID           uuid.NullUUID `bun:"type:uuid" json:"humanId"`
	Role              *Role         `bun:"rel:belongs-to" json:"role"`
	RoleID            uuid.NullUUID `bun:"type:uuid" json:"roleId"`
	Questions         Questions     `bun:"rel:has-many" json:"questions"`
	Comments          Comments      `bun:"rel:has-many" json:"comments"`
	RejectEmail       bool          `json:"rejectEmail"`
	CreatedAt         time.Time     `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	Children          Children      `bun:"rel:has-many" json:"children"`
	ChildrenForDelete []uuid.UUID   `bun:"-" json:"childrenForDelete"`

	DonorRulesUsers          DonorRulesUsers `bun:"rel:has-many" json:"donorRulesUsers"`
	DoctorsUsers             DoctorsUsers    `bun:"rel:has-many" json:"doctorsUsers"`
	DonorRulesUsersForDelete []uuid.UUID     `bun:"-" json:"donorRulesUsersForDelete"`

	DpoApplications          DpoApplications `bun:"rel:has-many" json:"dpoApplications"`
	DpoApplicationsForDelete []uuid.UUID     `bun:"-" json:"dpoApplicationsForDelete"`

	PostgraduateApplications          PostgraduateApplications `bun:"rel:has-many" json:"postgraduateApplications"`
	PostgraduateApplicationsForDelete []uuid.UUID              `bun:"-" json:"postgraduateApplicationsForDelete"`

	CandidateApplications          CandidateApplications `bun:"rel:has-many" json:"candidateApplications"`
	CandidateApplicationsForDelete []uuid.UUID           `bun:"-" json:"сandidateApplicationsForDelete"`

	ResidencyApplications ResidencyApplications `bun:"rel:has-many,join:id=user_id" json:"residencyApplications"`
	// CandidateApplicationsForDelete []uuid.UUID           `bun:"-" json:"сandidateApplicationsForDelete"`

	VisitsApplications          VisitsApplications `bun:"rel:has-many" json:"visitsApplications"`
	VisitsApplicationsForDelete []uuid.UUID        `bun:"-" json:"visitsApplicationsForDelete"`

	VacancyResponses          VacancyResponses `bun:"rel:has-many" json:"vacancyResponses"`
	VacancyResponsesForDelete []uuid.UUID      `bun:"-" json:"vacancyResponsesForDelete"`

	DailyMenuOrders          DailyMenuOrders `bun:"rel:has-many,join:id=user_id" json:"dailyMenuOrders"`
	DailyMenuOrdersForDelete []uuid.UUID     `bun:"-" json:"dailyMenuOrdersForDelete"`

	// FormValues FormValues `bun:"rel:has-many" json:"formValues"`
	UserAccountID uuid.NullUUID `bun:"type:uuid" json:"-"`
}

type Users []*User

func (i *User) GenerateHashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(i.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	i.Password = string(hash)
	return nil
}

func (i *User) CompareWithUUID(externalUUID string) bool {
	return i.UUID.String() == externalUUID
}

func (i *User) CompareWithHashPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(i.Password), []byte(password)) == nil
}

func (i *User) SetForeignKeys() {
	i.HumanID = i.Human.ID
	if i.Role != nil && i.Role.ID.Valid {
		i.RoleID = i.Role.ID
	}
}

func (i *User) SetFilePath(fileID *string) *string {
	path := i.Human.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (i *User) SetIDForChildren() {
	for index := range i.Children {
		i.Children[index].UserID = i.ID
	}
	if len(i.DonorRulesUsers) > 0 {
		for index := range i.Children {
			i.DonorRulesUsers[index].UserID = i.ID.UUID
		}
	}
}

func (i *User) SetJWTClaimsMap(claims map[string]interface{}) {
	claims[middleware.ClaimUserID.String()] = i.ID.UUID
}
