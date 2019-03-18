package users

type mockRepository []User

// NewMockRepository returns new mockRepository repository.
func NewMockRepository() Repository {
	return mockRepository{
		User{
			Email:    "user1@gmail.com",
			Name:     "User Friendly",
			Role:     "user",
			ID:       42,
			Password: "user1",
		},
		User{
			Email:    "user2@gmail.com",
			Name:     "John Walker",
			Role:     "user",
			ID:       45,
			Password: "12345qwerty",
		},
	}
}

// GetByCreds returns user with provided credentials
// or returns ErrNoSuchUser if it does not exist.
func (repo mockRepository) GetByCreds(email, password string) (*User, error) {
	for _, u := range repo {
		if u.Email == email && u.Password == password {
			return &u, nil
		}
	}

	return nil, ErrNoSuchUser
}

// GetByID returns user with provided id or
// returns ErrNoSuchUser if it does not exist.
func (repo mockRepository) GetByID(id int) (*User, error) {
	for _, u := range repo {
		if u.ID == id {
			return &u, nil
		}
	}

	return nil, ErrNoSuchUser
}
