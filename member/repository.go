package member

import "gorm.io/gorm"

type Repository interface {
	Index() ([]Member, error)
	Store(member Member) (Member, error)
	Show(ID int) (Member, error)
	Update(member Member) (Member, error)
	Destroy(member Member) (Member, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Index() ([]Member, error) {
	var member []Member

	err := r.db.Find(&member).Error
	if err != nil {
		return member, err
	}

	return member, nil
}

func (r *repository) Store(member Member) (Member, error) {
	err := r.db.Create(&member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}

func (r *repository) Show(ID int) (Member, error) {
	var member Member
	err := r.db.Where("id = ? ", ID).Find(&member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}

func (r *repository) Update(member Member) (Member, error) {
	err := r.db.Save(&member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}

func (r *repository) Destroy(member Member) (Member, error) {
	err := r.db.Delete(&member).Error
	if err != nil {
		return member, err
	}
	return member, nil
}
