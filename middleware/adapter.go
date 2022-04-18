package middleware

import (
	"context"
	"fmt"
	//"github.com/casbin/casbin/v2/model"
	//"github.com/casbin/casbin/v2/persist"
	"github.com/google/uuid"
	"github.com/mmcloughlin/meow"
	"github.com/uptrace/bun"
	"strings"
)

const DefaultTableName = "casbin_rules"

// CasbinRule represents a rule in Casbin.
type CasbinRule struct {
	bun.BaseModel `bun:"casbin_rules,alias:casbin_rule"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Ptype         string
	V0            string
	V1            string `bun:unique`
	V2            string
	V3            string
	V4            string
	V5            string
}

func (r *CasbinRule) String() string {
	const prefixLine = ", "
	var sb strings.Builder

	sb.Grow(
		len(r.Ptype) +
			len(r.V0) + len(r.V1) + len(r.V2) +
			len(r.V3) + len(r.V4) + len(r.V5),
	)

	sb.WriteString(r.Ptype)
	if len(r.V0) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(r.V0)
	}
	if len(r.V1) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(r.V1)
	}
	if len(r.V2) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(r.V2)
	}
	if len(r.V3) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(r.V3)
	}
	if len(r.V4) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(r.V4)
	}
	if len(r.V5) > 0 {
		sb.WriteString(prefixLine)
		sb.WriteString(r.V5)
	}

	return sb.String()
}

type Filter struct {
	P []string
	G []string
}

// Adapter represents the github.com/go-pg/pg adapter for policy storage.
type Adapter struct {
	db              *bun.DB
	tableName       string
	skipTableCreate bool
	filtered        bool
}

type Option func(a *Adapter)

// NewAdapterByDB creates new Adapter by using existing DB connection
// creates table from CasbinRule struct if it doesn't exist
func NewAdapterByDB(db *bun.DB, opts ...Option) (*Adapter, error) {
	a := &Adapter{db: db, tableName: DefaultTableName}
	for _, opt := range opts {
		opt(a)
	}

	if !a.skipTableCreate {
		if err := a.createTableifNotExists(); err != nil {
			return nil, fmt.Errorf("pgadapter.NewAdapter: %v", err)
		}
	}
	return a, nil
}

func (a *Adapter) createTableifNotExists() error {
	fmt.Println(a.tableName)
	_, err := a.db.NewCreateTable().Model((*CasbinRule)(nil)).Table(a.tableName).IfNotExists().Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.db.NewInsert().Model(line).
		Table(a.tableName).
		//OnConflict("DO NOTHING").
		Exec(context.Background())

	return err

	return err
}

// LoadPolicy loads policy from database.
//func (a *Adapter) LoadPolicy(model model.Model) error {
//	var lines []*CasbinRule
//
//	err := a.db.NewSelect().Model(&lines).Table(a.tableName).Scan(context.Background())
//	if err != nil {
//		return err
//	}
//	for _, line := range lines {
//		persist.LoadPolicyLine(line.String(), model)
//	}
//
//	a.filtered = false
//
//	return nil
//}

func savePolicyLine(ptype string, rule []string) *CasbinRule {
	line := &CasbinRule{Ptype: ptype}

	l := len(rule)
	if l > 0 {
		line.V0 = rule[0]
	}
	if l > 1 {
		line.V1 = rule[1]
	}
	if l > 2 {
		line.V2 = rule[2]
	}
	if l > 3 {
		line.V3 = rule[3]
	}
	if l > 4 {
		line.V4 = rule[4]
	}
	if l > 5 {
		line.V5 = rule[5]
	}
	//res, _ := strconv.Atoi(policyID(ptype, rule))
	//line.ID = int64(res)

	return line
}

func policyID(ptype string, rule []string) string {
	data := strings.Join(append([]string{ptype}, rule...), ",")
	sum := meow.Checksum(0, []byte(data))
	return fmt.Sprintf("%x", sum)
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	query := a.db.NewDelete().Model((*CasbinRule)(nil)).Table(a.tableName).Where("ptype = ?", ptype)

	idx := fieldIndex + len(fieldValues)
	if fieldIndex <= 0 && idx > 0 && fieldValues[0-fieldIndex] != "" {
		query = query.Where("v0 = ?", fieldValues[0-fieldIndex])
	}
	if fieldIndex <= 1 && idx > 1 && fieldValues[1-fieldIndex] != "" {
		query = query.Where("v1 = ?", fieldValues[1-fieldIndex])
	}
	if fieldIndex <= 2 && idx > 2 && fieldValues[2-fieldIndex] != "" {
		query = query.Where("v2 = ?", fieldValues[2-fieldIndex])
	}
	if fieldIndex <= 3 && idx > 3 && fieldValues[3-fieldIndex] != "" {
		query = query.Where("v3 = ?", fieldValues[3-fieldIndex])
	}
	if fieldIndex <= 4 && idx > 4 && fieldValues[4-fieldIndex] != "" {
		query = query.Where("v4 = ?", fieldValues[4-fieldIndex])
	}
	if fieldIndex <= 5 && idx > 5 && fieldValues[5-fieldIndex] != "" {
		query = query.Where("v5 = ?", fieldValues[5-fieldIndex])
	}

	_, err := query.Exec(context.Background())
	return err

}

// RemovePolicy removes a policy rule from the storage.
func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	line := savePolicyLine(ptype, rule)
	_, err := a.db.NewDelete().Model(line).Table(a.tableName).WherePK().Exec(context.Background())
	return err
}

//SavePolicy saves policy to database.
//func (a *Adapter) SavePolicy(model model.Model) error {
//
//	_, err := a.db.NewDelete().Table(a.tableName).Where("id IS NOT NULL").Exec(context.Background())
//	if err != nil {
//		return err
//	}
//
//	var lines []*CasbinRule
//
//	for ptype, ast := range model["p"] {
//		for _, rule := range ast.Policy {
//			line := savePolicyLine(ptype, rule)
//			lines = append(lines, line)
//		}
//	}
//
//	for ptype, ast := range model["g"] {
//		for _, rule := range ast.Policy {
//			line := savePolicyLine(ptype, rule)
//			lines = append(lines, line)
//		}
//	}
//
//	for _, line := range lines {
//		_, err = a.db.NewInsert().Model(line).Table(a.tableName).Exec(context.Background())
//		if err != nil {
//			return err
//		}
//	}
//
//	if err != nil {
//		return fmt.Errorf("commit DB transaction: %v", err)
//	}
//
//	return nil
//}
