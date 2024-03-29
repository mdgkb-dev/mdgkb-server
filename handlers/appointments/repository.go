package appointments

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.Appointments, error) {
	items := make(models.Appointments, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("Doctor.Employee.Human").
		Relation("Doctor.DoctorsDivisions").
		Relation("FormValue.Child.Human").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus")

	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Appointment, error) {
	item := models.Appointment{}
	err := r.db().NewSelect().Model(&item).
		Relation("Doctor.Employee.Human").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormValue.Fields.File").
		Relation("FormValue.FormValueFiles.File").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("AppointmentType.FormPattern", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.ExcludeColumn("with_personal_data_agreement")
		}).
		Relation("AppointmentType.FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("AppointmentType.FormPattern.Fields.ValueType").
		Where("appointments_view.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.Appointment) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Appointment{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Appointment) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Appointment)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Appointments) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		//Set("doctor_id = EXCLUDED.doctor_id").
		//Set("position = EXCLUDED.position").
		Model(&items).
		Exec(r.ctx)
	return err
}
