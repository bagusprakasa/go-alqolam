package wallet

import "gorm.io/gorm"

type Repository interface {
	Index() ([]Wallet, error)
	Store(awllet Wallet) (Wallet, error)
	Show(ID int) (Wallet, error)
	Update(awllet Wallet) (Wallet, error)
	Destroy(awllet Wallet) (Wallet, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Index() ([]Wallet, error) {
	var wallet []Wallet

	err := r.db.Find(&wallet).Error
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (r *repository) Store(wallet Wallet) (Wallet, error) {
	err := r.db.Create(&wallet).Error
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (r *repository) Show(ID int) (Wallet, error) {
	var wallet Wallet
	err := r.db.Where("id = ? ", ID).Find(&wallet).Error
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (r *repository) Update(wallet Wallet) (Wallet, error) {
	err := r.db.Save(&wallet).Error
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (r *repository) Destroy(wallet Wallet) (Wallet, error) {
	err := r.db.Delete(&wallet).Error
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}
