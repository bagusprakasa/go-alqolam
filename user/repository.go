package user

import "gorm.io/gorm"

type Repository interface {
	Index() ([]User, error)
	Store(user User) (User, error)
	Show(ID int) (User, error)
	Update(user User) (User, error)
	Destroy(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Index() ([]User, error) {
	var user []User

	err := r.db.Where("deleted_at = ?", "0000-00-00 00:00:00.000").Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Store(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Show(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ? ", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ? ", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Destroy(user User) (User, error) {
	err := r.db.Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
