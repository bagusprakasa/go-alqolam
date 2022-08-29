package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Index() ([]User, error)
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	Show(ID int) (User, error)
	FindUser(ID int) (User, error)
	Update(inputID GetDetailInput, inputData UpdatedUserInput) (User, error)
	Destroy(inputID GetDetailInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Index() ([]User, error) {

	users, err := s.repository.Index()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(password)
	user.Role = "admin"

	newUser, err := s.repository.Store(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) Show(ID int) (User, error) {
	user, err := s.repository.Show(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) FindUser(ID int) (User, error) {
	user, err := s.repository.Show(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) Update(inputID GetDetailInput, inputData UpdatedUserInput) (User, error) {
	user, err := s.repository.Show(inputID.ID)
	if err != nil {
		return user, err
	}

	user.Name = inputData.Name
	user.Email = inputData.Email

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) Destroy(inputID GetDetailInput) (User, error) {
	user, err := s.repository.Show(inputID.ID)
	if err != nil {
		return user, err
	}
	deletedUser, err := s.repository.Destroy(user)
	if err != nil {
		return deletedUser, err
	}

	return deletedUser, nil
}
