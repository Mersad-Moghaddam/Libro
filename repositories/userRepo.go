package repositories

import (
	"libro/models/user"

	"gorm.io/gorm"
)

type UserRepo struct{ db *gorm.DB }

func NewUserRepo(db *gorm.DB) *UserRepo { return &UserRepo{db: db} }

func (r *UserRepo) Create(u *user.User) error { return r.db.Create(u).Error }
func (r *UserRepo) FindByEmail(email string) (*user.User, error) {
	var u user.User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *UserRepo) FindByID(id uint) (*user.User, error) {
	var u user.User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *UserRepo) Update(u *user.User) error { return r.db.Save(u).Error }
