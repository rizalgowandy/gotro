package rqAuth

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/kokizzu/gotro/L"
)

func (s *Users) AllOffsetLimit(offset, limit uint32) (res []*Users) {
	rows := s.FindOffsetLimit(offset, limit, s.UniqueIndexId())
	res = make([]*Users, 0, len(rows))
	for i := range rows {
		rows[i].CensorFields()
		res = append(res, &rows[i])
	}
	return
}

func (s *Users) CensorFields() {
	s.Password = ``
	s.SecretCode = ``
}

func (s *Users) CheckPassword(currentPassword string) bool {
	hash := []byte(s.Password)
	pass := []byte(currentPassword)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return !L.IsError(err, `bcrypt.CompareHashAndPassword`)
}

// add more custom methods here
