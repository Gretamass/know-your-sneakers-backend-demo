package db

import "github.com/Gretamass/know-your-sneakers-backend/internal/user"

// GetUser retrieves a user by username
func GetUser(username string) (*user.User, error) {
	// query the database for the user
	// ...

	// scan the result into a user struct
	var u user.User
	// ...

	return &u, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id int64) (*user.User, error) {
	// query the database for the user
	// ...

	// scan the result into a user struct
	var u user.User
	// ...

	return &u, nil
}

// CreateUser creates a new user in the database
func CreateUser(u *user.User) error {
	// insert the user into the database
	// ...

	return nil
}

// IsUsernameTaken checks if a username is already taken
func IsUsernameTaken(username string) bool {
	// query the database for the user
	// ...

	// check if the user exists
	// ...

	return false
}
