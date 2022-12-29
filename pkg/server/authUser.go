package server

// User contains user's information
type User struct {
	Name  string
	Email string
}

// NewUser returns a new user
func NewUserAuth(name string, email string) (*User, error) {
	user := &User{
		Name:  name,
		Email: email,
	}

	return user, nil
}
