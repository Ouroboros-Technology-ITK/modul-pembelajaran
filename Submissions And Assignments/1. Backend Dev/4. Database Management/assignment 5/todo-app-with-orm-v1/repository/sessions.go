package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	return nil // TODO: replace this
}

func (u *SessionsRepository) DeleteSession(token string) error {
	return nil // TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	return nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	return model.Session{}, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	return model.Session{}, nil // TODO: replace this
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	session := model.Session{} // TODO: replace this

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}

	return session, nil
}

func (u *SessionsRepository) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
