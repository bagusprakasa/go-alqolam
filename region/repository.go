package region

import "gorm.io/gorm"

type Repository interface {
	Index() ([]Region, error)
	Store(region Region) (Region, error)
	Show(ID int) (Region, error)
	Update(region Region) (Region, error)
	Destroy(region Region) (Region, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Index() ([]Region, error) {
	var region []Region

	err := r.db.Find(&region).Error
	if err != nil {
		return region, err
	}

	return region, nil
}

func (r *repository) Store(region Region) (Region, error) {
	err := r.db.Create(&region).Error
	if err != nil {
		return region, err
	}
	return region, nil
}

func (r *repository) Show(ID int) (Region, error) {
	var region Region
	err := r.db.Where("id = ? ", ID).Find(&region).Error
	if err != nil {
		return region, err
	}
	return region, nil
}

func (r *repository) Update(region Region) (Region, error) {
	err := r.db.Save(&region).Error
	if err != nil {
		return region, err
	}
	return region, nil
}

func (r *repository) Destroy(region Region) (Region, error) {
	err := r.db.Delete(&region).Error
	if err != nil {
		return region, err
	}
	return region, nil
}
