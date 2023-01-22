package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	Id        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var todos = []todo{
	{Id: "1", Item: "Clean room", Completed: false},
	{Id: "2", Item: "Meal prep", Completed: false},
	{Id: "3", Item: "Record video", Completed: false},
}

var users = []user{
	{Email: "johndoe@example.com", Password: "johndoe"},
	{Email: "janesmith@example.com", Password: "janesmith"},
	{Email: "bobjohnson@example.com", Password: "bobjohnson"},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/users", getUsers)

	router.POST("/todos", addTodo)
	router.POST("/user", addUser)

	router.GET("/todos/:id", getTodo)

	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:9090")

}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func addUser(context *gin.Context) {
	var newUser user

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}
