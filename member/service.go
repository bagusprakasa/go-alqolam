package member

type Service interface {
	Index() ([]Member, error)
	Store(inputData MemberInput) (Member, error)
	Show(ID int) (Member, error)
	Update(inputID GetDetailInput, inputData MemberInput) (Member, error)
	Destroy(inputID GetDetailInput) (Member, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]Member, error) {

	members, err := s.repository.Index()
	if err != nil {
		return members, err
	}

	return members, nil
}

func (s *service) Store(input MemberInput) (Member, error) {
	member := Member{}
	member.RegionID = input.RegionID
	member.Name = input.Name
	member.Phone = input.Phone
	member.Address = input.Address
	member.Gender = input.Gender
	newMember, err := s.repository.Store(member)
	if err != nil {
		return newMember, err
	}

	return newMember, nil
}

func (s *service) Show(ID int) (Member, error) {
	member, err := s.repository.Show(ID)
	if err != nil {
		return member, err
	}

	return member, nil
}

func (s *service) Update(inputID GetDetailInput, inputData MemberInput) (Member, error) {
	member, err := s.repository.Show(inputID.ID)
	if err != nil {
		return member, err
	}

	member.RegionID = inputData.RegionID
	member.Name = inputData.Name
	member.Phone = inputData.Phone
	member.Address = inputData.Address
	member.Gender = inputData.Gender

	updatedMember, err := s.repository.Update(member)
	if err != nil {
		return updatedMember, err
	}

	return updatedMember, nil
}

func (s *service) Destroy(inputID GetDetailInput) (Member, error) {
	member, err := s.repository.Show(inputID.ID)
	if err != nil {
		return member, err
	}
	deletedMember, err := s.repository.Destroy(member)
	if err != nil {
		return deletedMember, err
	}

	return deletedMember, nil
}
