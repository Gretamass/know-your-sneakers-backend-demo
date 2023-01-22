package user

import (
	"net/http"

	"github.com/Gretamass/know-your-sneakers-backend/internal/db"
	"github.com/Gretamass/know-your-sneakers-backend/pkg/errors"
)

// User contains information about a user
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginHandler handles login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// parse the request body to get the username and password
	// ...

	// check if the user exists in the database
	user, err := db.GetUser(username)
	if err != nil {
		errors.WriteError(w, errors.ErrInvalidCredentials)
		return
	}

	// compare the password from the request with the password from the database
	if user.Password != password {
		errors.WriteError(w, errors.ErrInvalidCredentials)
		return
	}

	// if the login is successful, set a session cookie
	// ...

	// write a JSON response indicating that the login was successful
	// ...
}

// LogoutHandler handles logout requests
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// remove the session cookie
	// ...

	// write a JSON response indicating that the logout was successful
	// ...
}

// SignupHandler handles signup requests
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// parse the request body to get the username and password
	// ...

	// check if the username is already taken
	if db.IsUsernameTaken(username) {
		errors.WriteError(w, errors.ErrUsernameTaken)
		return
	}

	// create a new user in the database
	user := &User{
		Username: username,
		Password: password,
	}
	if err := db.CreateUser(user); err != nil {
		errors.WriteError(w, errors.ErrServerError)
		return
	}

	// if the signup is successful, set a session cookie
	// ...

	// write a JSON response indicating that the signup was successful
	// ...
}

// AccountHandler handles requests for the user's account page
func AccountHandler(w http.ResponseWriter, r *http.Request) {
	// get the user's ID from the session cookie
	// ...

	// get the user's information from the database
	user, err := db.GetUserByID(userID)
	if err != nil {
		errors.WriteError(w, errors.ErrServerError)
		return
	}

	// render the account template and pass the user's information to it
	// ...
}
