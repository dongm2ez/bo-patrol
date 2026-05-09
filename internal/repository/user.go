package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/bo-patrol/internal/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) List(page, pageSize int) ([]domain.User, int64, error) {
	var users []domain.User
	var total int64

	offset := (page - 1) * pageSize
	r.db.Model(&domain.User{}).Count(&total)
	err := r.db.Offset(offset).Limit(pageSize).Find(&users).Error

	return users, total, err
}

func (r *UserRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&domain.User{}).Where("id = ?", id).Updates(updates).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *UserRepository) VerifyPassword(user *domain.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
