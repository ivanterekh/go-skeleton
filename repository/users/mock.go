package users

import "github.com/ivanterekh/go-skeleton/model"

type mock []model.User

// NewMock returns new mock repository.
func NewMock() Repository {
	return mock{
		model.User{
			Email:    "user1@gmail.com",
			Name:     "User Friendly",
			Role:     "user",
			ID:       42,
			Password: "user1",
		},
		model.User{
			Email:    "user2@gmail.com",
			Name:     "John Walker",
			Role:     "user",
			ID:       45,
			Password: "12345qwerty",
		},
	}
}

// ByCreds returns user with provided credentials
// or returns ErrNoSuchUser.
func (repo mock) ByCreds(email, password string) (*model.User, error) {
	for _, u := range repo {
		if u.Email == email && u.Password == password {
			return &u, nil
		}
	}

	return nil, ErrNoSuchUser
}

func (repo mock) ByID(id int) (*model.User, error) {
	for _, u := range repo {
		if u.ID == id {
			return &u, nil
		}
	}

	return nil, ErrNoSuchUser
}
