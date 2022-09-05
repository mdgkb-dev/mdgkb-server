package tests

import (
	"strconv"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
)

var funcMap = template.FuncMap{
	"uuid": func() string {
		return CreateUUID.UUID.String()
	},
}

func GetFixtures(db *bun.DB, models ...interface{}) *dbfixture.Fixture {
	db.RegisterModel(models...)
	return dbfixture.New(db, dbfixture.WithTemplateFuncs(funcMap), dbfixture.WithTruncateTables())
}

// Numbers mocks
var (
	CreateUUID        = uuid.NullUUID{UUID: uuid.MustParse("216165b0-e20c-4849-b25d-1d65e0a1700e"), Valid: true}
	CreateUintID      = uint(1001)
	ForeignUintID     = uint(101)
	ForeignUUID       = uint(101)
	DeleteUintID      = uint(10)
	NotExistingUintID = uint(14999999)
	ZeroUintID        = uint(0)
	CorrectUint8      = uint8(0)
	CorrectFloat      = 1.02
	CorrectFloat32    = float32(1.02)
	CorrectBool       = true
)

// PAGINATIONS
//
//// TruePagination mock
//var TruePagination = model.Pagination{
//	ID:    CreateUintID,
//	Limit: int(ForeignUintID),
//	From:  int(ForeignUintID),
//	Name:  "trading_code_id",
//	Sort:  "ASC",
//}

// BatchTestName mocks
var (
	BatchTestName0 = "BatchName0"
	BatchTestName1 = "BatchName1"
	BatchTestName2 = "BatchName2"
)

// STRINGS

// TestString mock
var (
	TestString          = "TestString" // всегда проверять согласованность этой переменной с переменной в test_seeding.sql
	CreateStringID      = strconv.Itoa(int(CreateUintID))
	ForeignStringID     = strconv.Itoa(int(ForeignUintID))
	ZeroStringID        = strconv.Itoa(int(ZeroUintID))
	NotExistingStringID = strconv.Itoa(int(NotExistingUintID))
	NotExistingString   = "falgm3v67opt4!%$2"
	EmptyString         = ""
	ExistingInn         = "0000000001"
	ConstPassword       = "11111111111111111111111111111111"
)

// Months mocks
var (
	CorrectMonth   = "2020-02"
	LettersMonth   = "asdsafg"
	EmptyMonth     = ""
	IncorrectMonth = "1242153-215"
	LessZeroMonth  = "-2020-05"
)

// Dates mocks
var (
	NowDate        = time.Now()
	PastData       = time.Date(2000, 01, 01, 1, 1, 1, 1, time.UTC)
	NotCorrectData = time.Date(2000001, 0231, -001, 11313, 6731, 2461, -13251, time.UTC)
)

// BOOLS

var TrueMock = true

// MockEmail
type MockEmail struct{}

// SetRequest struct
func (e *MockEmail) SetRequest(from string, to []string, subject string) {
}

// GetFromUser struct
func (e *MockEmail) GetFromUser() string {
	return ""
}

// SendHTMLEmail func
func (e *MockEmail) SendHTMLEmail(templateData interface{}, templates ...string) error {
	return nil
}

// SendEmail func
func (e *MockEmail) SendEmail() error {
	return nil
}

// ParseTemplate func
func (e *MockEmail) ParseTemplate(data interface{}, templates ...string) error {
	return nil
}
