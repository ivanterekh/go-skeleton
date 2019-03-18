package users

// User represents a service user.
type User struct {
	ID       int
	Role     string
	Name     string
	Email    string
	Password string
}
