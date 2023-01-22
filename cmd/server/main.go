package main

import (
	"net/http"
)

func main() {
	// create a new router
	router := http.NewServeMux()

	// add middleware to the router
	//router.Handle("/", middleware.Chain(sneaker.HomeHandler, middleware.Logging()))
	//router.Handle("/sneakers", middleware.Chain(sneaker.ListHandler, middleware.Auth()))
	//router.Handle("/sneakers/{id}", middleware.Chain(sneaker.DetailHandler, middleware.Auth()))
	//router.Handle("/account", middleware.Chain(user.AccountHandler, middleware.Auth()))
	//router.Handle("/login", user.LoginHandler)
	//router.Handle("/logout", middleware.Chain(user.LogoutHandler, middleware.Auth()))
	//router.Handle("/signup", user.SignupHandler)

	// start the server
	http.ListenAndServe(":8080", router)
}
