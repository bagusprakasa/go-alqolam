package transferwallet

type Service interface {
	Index() ([]TransferWallet, error)
	Store(input TransferWalletInput) (TransferWallet, error)
	Show(ID int) (TransferWallet, error)
	Update(inputID GetDetailInput, inputData TransferWalletInput) (TransferWallet, error)
	Destroy(inputID GetDetailInput) (TransferWallet, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]TransferWallet, error) {

	transferWallets, err := s.repository.Index()
	if err != nil {
		return transferWallets, err
	}

	return transferWallets, nil
}

func (s *service) Store(input TransferWalletInput) (TransferWallet, error) {
	transferWallet := TransferWallet{}
	transferWallet.FromWalletId = input.FromWalletId
	transferWallet.ToWalletId = input.ToWalletId
	transferWallet.Total = input.Total
	transferWallet.Date = input.Date

	transferWalletNew, err := s.repository.Store(transferWallet)
	if err != nil {
		return transferWalletNew, err
	}

	return transferWalletNew, nil
}

func (s *service) Show(ID int) (TransferWallet, error) {
	transferWallet, err := s.repository.Show(ID)
	if err != nil {
		return transferWallet, err
	}

	return transferWallet, nil
}

func (s *service) Update(inputID GetDetailInput, inputData TransferWalletInput) (TransferWallet, error) {
	transferWallet, err := s.repository.Show(inputID.ID)
	if err != nil {
		return transferWallet, err
	}

	transferWallet.FromWalletId = inputData.FromWalletId
	transferWallet.ToWalletId = inputData.ToWalletId
	transferWallet.Total = inputData.Total
	transferWallet.Date = inputData.Date

	updatedTransferWallet, err := s.repository.Update(transferWallet)
	if err != nil {
		return updatedTransferWallet, err
	}

	return updatedTransferWallet, nil
}

func (s *service) Destroy(inputID GetDetailInput) (TransferWallet, error) {
	transferWallet, err := s.repository.Show(inputID.ID)
	if err != nil {
		return transferWallet, err
	}
	deletedTransferWallet, err := s.repository.Destroy(transferWallet)
	if err != nil {
		return deletedTransferWallet, err
	}

	return deletedTransferWallet, nil
}
