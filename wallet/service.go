package wallet

type Service interface {
	Index() ([]Wallet, error)
	Store(input WalletInput) (Wallet, error)
	Show(ID int) (Wallet, error)
	Update(inputID GetDetailInput, inputData WalletInput) (Wallet, error)
	Destroy(inputID GetDetailInput) (Wallet, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]Wallet, error) {

	wallets, err := s.repository.Index()
	if err != nil {
		return wallets, err
	}

	return wallets, nil
}

func (s *service) Store(input WalletInput) (Wallet, error) {
	wallet := Wallet{}
	wallet.Type = input.Type
	wallet.Name = input.Name
	wallet.AccountNumber = input.AccountNumber
	wallet.Balance = 0

	newWaller, err := s.repository.Store(wallet)
	if err != nil {
		return newWaller, err
	}

	return newWaller, nil
}

func (s *service) Show(ID int) (Wallet, error) {
	wallet, err := s.repository.Show(ID)
	if err != nil {
		return wallet, err
	}

	return wallet, nil
}

func (s *service) Update(inputID GetDetailInput, inputData WalletInput) (Wallet, error) {
	wallet, err := s.repository.Show(inputID.ID)
	if err != nil {
		return wallet, err
	}

	wallet.Type = inputData.Type
	wallet.Name = inputData.Name
	wallet.AccountNumber = inputData.AccountNumber

	updatedWallet, err := s.repository.Update(wallet)
	if err != nil {
		return updatedWallet, err
	}

	return updatedWallet, nil
}

func (s *service) Destroy(inputID GetDetailInput) (Wallet, error) {
	wallet, err := s.repository.Show(inputID.ID)
	if err != nil {
		return wallet, err
	}
	deletedWallet, err := s.repository.Destroy(wallet)
	if err != nil {
		return deletedWallet, err
	}

	return deletedWallet, nil
}
