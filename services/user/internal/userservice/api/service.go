package api

import (
	"rideshare/services/user/internal/userservice/internal/db"
	"rideshare/services/user/internal/userservice/internal/publisher"

	"golang.org/x/crypto/bcrypt"
)

type userRepository interface {
	CreateUser(us db.User) error
	GetUserByID(userID int) (db.User, error)
}

type publisherI interface {
	Publish(string) error
}

type userService struct {
	repo      userRepository
	publisher publisherI
}

type ServiceConfig struct {
}

func NewUserService(cfg ServiceConfig) (*userService, error) {

	userRepo := db.NewInMemoryUserRepository()

	p, err := publisher.NewPublisher(publisher.Config{
		Bootstrap: "localhost:9092",
		Topic:     "users",
	})

	if err != nil {
		return nil, err
	}

	return &userService{repo: userRepo, publisher: p}, nil
}

func (s *userService) Register(param UserRegistrationParam) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := db.User{
		Name:           param.Name,
		Email:          param.Email,
		HashedPassword: string(hashedPassword),
	}
	err = s.repo.CreateUser(user)
	if err != nil {
		return err
	}

	err = s.publisher.Publish("user created")
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetUserByID(id int) (PublicUser, error) {
	u, err := s.repo.GetUserByID(id)
	if err != nil {
		return PublicUser{}, err // Return empty user and the error
	}
	return NewPublicUserFromInternal(u), nil
}
