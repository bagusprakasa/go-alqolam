package transferwallet

import "gorm.io/gorm"

type Repository interface {
	Index() ([]TransferWallet, error)
	Store(transferWallet TransferWallet) (TransferWallet, error)
	Show(ID int) (TransferWallet, error)
	Update(transferWallet TransferWallet) (TransferWallet, error)
	Destroy(transferWallet TransferWallet) (TransferWallet, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Index() ([]TransferWallet, error) {
	var transferWallet []TransferWallet

	err := r.db.Preload("FromWallet").Preload("ToWallet").Find(&transferWallet).Error
	if err != nil {
		return transferWallet, err
	}

	return transferWallet, nil
}

func (r *repository) Store(transferWallet TransferWallet) (TransferWallet, error) {
	err := r.db.Create(&transferWallet).Error
	if err != nil {
		return transferWallet, err
	}
	return transferWallet, nil
}

func (r *repository) Show(ID int) (TransferWallet, error) {
	var transferWallet TransferWallet
	err := r.db.Preload("FromWallet").Preload("ToWallet").Where("id = ? ", ID).Find(&transferWallet).Error
	if err != nil {
		return transferWallet, err
	}
	return transferWallet, nil
}

func (r *repository) Update(transferWallet TransferWallet) (TransferWallet, error) {
	err := r.db.Save(&transferWallet).Error
	if err != nil {
		return transferWallet, err
	}
	return transferWallet, nil
}

func (r *repository) Destroy(transferWallet TransferWallet) (TransferWallet, error) {
	err := r.db.Delete(&transferWallet).Error
	if err != nil {
		return transferWallet, err
	}
	return transferWallet, nil
}
