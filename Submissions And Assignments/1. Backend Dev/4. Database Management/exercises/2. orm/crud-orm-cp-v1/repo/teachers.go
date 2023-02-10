package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	return nil // TODO: replace this
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	return []model.Teacher{}, nil // TODO: replace this
}

func (t TeacherRepo) Update(id uint, name string) error {
	return nil // TODO: replace this
}

func (t TeacherRepo) Delete(id uint) error {
	return nil // TODO: replace this
}
