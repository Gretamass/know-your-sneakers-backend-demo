package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type users struct {
	id        string `json:"id"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	email     string `json:"email"`
}

func main() {
	db, err := sql.Open("mysql", "u351287655_root:pKYS;0830@tcp(sql861.main-hosting.eu)/u351287655_kys")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success login into database!")

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var users users
		err = results.Scan(&users.id, &users.firstName, &users.lastName, &users.email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(users)
	}
}

//
//func addUser() {
//	insert, err := db.Query("INSERT INTO users VALUES('3', 'Gedas', 'Piliukas', 'tgp3@gmail.com')")
//	if err != nil {
//		panic(err.Error())
//	}
//	defer insert.Close()
//	fmt.Println("Yay, values added!")
//}
//
//func getAllUsers(db) {
//	results, err := db.Query("SELECT * FROM users")
//	if err != nil {
//		panic(err.Error())
//	}
//	for results.Next() {
//		var users users
//		err = results.Scan(&users)
//		if err != nil {
//			panic(err.Error())
//		}
//		fmt.Println(users)
//	}
//}
