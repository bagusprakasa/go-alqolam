package region

type Service interface {
	Index() ([]Region, error)
	Store(inputData RegionInput) (Region, error)
	Show(ID int) (Region, error)
	Update(inputID GetDetailInput, inputData RegionInput) (Region, error)
	Destroy(inputID GetDetailInput) (Region, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]Region, error) {

	regions, err := s.repository.Index()
	if err != nil {
		return regions, err
	}

	return regions, nil
}

func (s *service) Store(input RegionInput) (Region, error) {
	region := Region{}
	region.Name = input.Name
	newRegion, err := s.repository.Store(region)
	if err != nil {
		return newRegion, err
	}

	return newRegion, nil
}

func (s *service) Show(ID int) (Region, error) {
	region, err := s.repository.Show(ID)
	if err != nil {
		return region, err
	}

	return region, nil
}

func (s *service) Update(inputID GetDetailInput, inputData RegionInput) (Region, error) {
	region, err := s.repository.Show(inputID.ID)
	if err != nil {
		return region, err
	}

	region.Name = inputData.Name

	updatedRegion, err := s.repository.Update(region)
	if err != nil {
		return updatedRegion, err
	}

	return updatedRegion, nil
}

func (s *service) Destroy(inputID GetDetailInput) (Region, error) {
	region, err := s.repository.Show(inputID.ID)
	if err != nil {
		return region, err
	}
	deletedRegion, err := s.repository.Destroy(region)
	if err != nil {
		return deletedRegion, err
	}

	return deletedRegion, nil
}
