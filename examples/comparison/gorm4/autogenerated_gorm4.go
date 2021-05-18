// Code generated by go-queryset. DO NOT EDIT.
package gorm4

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set UserQuerySet

// UserQuerySet is an queryset type for User
type UserQuerySet struct {
	db *gorm.DB
}

// NewUserQuerySet constructs new UserQuerySet
func NewUserQuerySet(db *gorm.DB) UserQuerySet {
	return UserQuerySet{
		db: db.Model(&User{}),
	}
}

func (qs UserQuerySet) w(db *gorm.DB) UserQuerySet {
	return NewUserQuerySet(db)
}

func (qs UserQuerySet) Select(fields ...UserDBSchemaField) UserQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *User) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *User) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) All(ret *[]User) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtEq(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtGt(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtGte(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtLt(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtLte(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) CreatedAtNe(createdAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Delete() error {
	return qs.db.Delete(User{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(User{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(User{})
	return db.RowsAffected, db.Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtEq(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtGt(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtGte(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtIsNotNull() UserQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtIsNull() UserQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtLt(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtLte(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) DeletedAtNe(deletedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) GetUpdater() UserUpdater {
	return NewUserUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDEq(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDGt(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDGte(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDIn(ID ...uint) UserQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDLt(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDLte(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDNe(ID uint) UserQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) IDNotIn(ID ...uint) UserQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Limit(limit int) UserQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) Offset(offset int) UserQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs UserQuerySet) One(ret *User) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByCreatedAt() UserQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByDeletedAt() UserQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByID() UserQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByRating is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByRating() UserQuerySet {
	return qs.w(qs.db.Order("rating ASC"))
}

// OrderAscByRatingMarks is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByRatingMarks() UserQuerySet {
	return qs.w(qs.db.Order("rating_marks ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderAscByUpdatedAt() UserQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByCreatedAt() UserQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByDeletedAt() UserQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByID() UserQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByRating is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByRating() UserQuerySet {
	return qs.w(qs.db.Order("rating DESC"))
}

// OrderDescByRatingMarks is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByRatingMarks() UserQuerySet {
	return qs.w(qs.db.Order("rating_marks DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) OrderDescByUpdatedAt() UserQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// RatingEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingEq(rating int) UserQuerySet {
	return qs.w(qs.db.Where("rating = ?", rating))
}

// RatingGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingGt(rating int) UserQuerySet {
	return qs.w(qs.db.Where("rating > ?", rating))
}

// RatingGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingGte(rating int) UserQuerySet {
	return qs.w(qs.db.Where("rating >= ?", rating))
}

// RatingIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingIn(rating ...int) UserQuerySet {
	if len(rating) == 0 {
		qs.db.AddError(errors.New("must at least pass one rating in RatingIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("rating IN (?)", rating))
}

// RatingLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingLt(rating int) UserQuerySet {
	return qs.w(qs.db.Where("rating < ?", rating))
}

// RatingLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingLte(rating int) UserQuerySet {
	return qs.w(qs.db.Where("rating <= ?", rating))
}

// RatingMarksEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingMarksEq(ratingMarks int) UserQuerySet {
	return qs.w(qs.db.Where("rating_marks = ?", ratingMarks))
}

// RatingMarksGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingMarksGt(ratingMarks int) UserQuerySet {
	return qs.w(qs.db.Where("rating_marks > ?", ratingMarks))
}

// RatingMarksGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingMarksGte(ratingMarks int) UserQuerySet {
	return qs.w(qs.db.Where("rating_marks >= ?", ratingMarks))
}

// RatingMarksIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingMarksIn(ratingMarks ...int) UserQuerySet {
	if len(ratingMarks) == 0 {
		qs.db.AddError(errors.New("must at least pass one ratingMarks in RatingMarksIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("rating_marks IN (?)", ratingMarks))
}

// RatingMarksLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingMarksLt(ratingMarks int) UserQuerySet {
	return qs.w(qs.db.Where("rating_marks < ?", ratingMarks))
}

// RatingMarksLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingMarksLte(ratingMarks int) UserQuerySet {
	return qs.w(qs.db.Where("rating_marks <= ?", ratingMarks))
}

// RatingMarksNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingMarksNe(ratingMarks int) UserQuerySet {
	return qs.w(qs.db.Where("rating_marks != ?", ratingMarks))
}

// RatingMarksNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingMarksNotIn(ratingMarks ...int) UserQuerySet {
	if len(ratingMarks) == 0 {
		qs.db.AddError(errors.New("must at least pass one ratingMarks in RatingMarksNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("rating_marks NOT IN (?)", ratingMarks))
}

// RatingNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingNe(rating int) UserQuerySet {
	return qs.w(qs.db.Where("rating != ?", rating))
}

// RatingNotIn is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) RatingNotIn(rating ...int) UserQuerySet {
	if len(rating) == 0 {
		qs.db.AddError(errors.New("must at least pass one rating in RatingNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("rating NOT IN (?)", rating))
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtEq(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtGt(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtGte(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtLt(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtLte(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs UserQuerySet) UpdatedAtNe(updatedAt time.Time) UserQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetCreatedAt(createdAt time.Time) UserUpdater {
	u.fields[string(UserDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetDeletedAt(deletedAt *time.Time) UserUpdater {
	u.fields[string(UserDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetID(ID uint) UserUpdater {
	u.fields[string(UserDBSchema.ID)] = ID
	return u
}

// SetRating is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetRating(rating int) UserUpdater {
	u.fields[string(UserDBSchema.Rating)] = rating
	return u
}

// SetRatingMarks is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetRatingMarks(ratingMarks int) UserUpdater {
	u.fields[string(UserDBSchema.RatingMarks)] = ratingMarks
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u UserUpdater) SetUpdatedAt(updatedAt time.Time) UserUpdater {
	u.fields[string(UserDBSchema.UpdatedAt)] = updatedAt
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u UserUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u UserUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set UserQuerySet

// ===== BEGIN of User modifiers

// UserDBSchemaField describes database schema field. It requires for method 'Update'
type UserDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f UserDBSchemaField) String() string {
	return string(f)
}

// UserDBSchema stores db field names of User
var UserDBSchema = struct {
	ID          UserDBSchemaField
	CreatedAt   UserDBSchemaField
	UpdatedAt   UserDBSchemaField
	DeletedAt   UserDBSchemaField
	Rating      UserDBSchemaField
	RatingMarks UserDBSchemaField
}{

	ID:          UserDBSchemaField("id"),
	CreatedAt:   UserDBSchemaField("created_at"),
	UpdatedAt:   UserDBSchemaField("updated_at"),
	DeletedAt:   UserDBSchemaField("deleted_at"),
	Rating:      UserDBSchemaField("rating"),
	RatingMarks: UserDBSchemaField("rating_marks"),
}

// Update updates User fields by primary key
// nolint: dupl
func (o *User) Update(db *gorm.DB, fields ...UserDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":           o.ID,
		"created_at":   o.CreatedAt,
		"updated_at":   o.UpdatedAt,
		"deleted_at":   o.DeletedAt,
		"rating":       o.Rating,
		"rating_marks": o.RatingMarks,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update User %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// UserUpdater is an User updates manager
type UserUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewUserUpdater creates new User updater
// nolint: dupl
func NewUserUpdater(db *gorm.DB) UserUpdater {
	return UserUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&User{}),
	}
}

// ===== END of User modifiers

// ===== END of all query sets
