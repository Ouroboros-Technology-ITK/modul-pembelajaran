package handler

import (
	"a21hc3NpZ25tZW50/app/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (u *TeacherRepo) AddTeacher(teacher model.Teacher) error {
	return nil // TODO: replace this
}

func (u *TeacherRepo) ReadTeacher() ([]model.ViewTeacher, error) {
	return nil,nil // TODO: replace this
}

func (u *TeacherRepo) UpdateName(id uint, name string) error {
	return nil // TODO: replace this
}

func (u *TeacherRepo) DeleteTeacher(id uint) error {
	return nil // TODO: replace this
}
