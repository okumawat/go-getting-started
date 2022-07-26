package main

import (
	"fmt"
	"net/http"

	"githum.com/okumawat/go-getting-started/contollers"
	"githum.com/okumawat/go-getting-started/models"
)

func main() {
	user := models.User{Name: "GoLang", Email: "test@go.com"}

	fmt.Println(user)

	//*models.Users = append(models.Users, user)
	models.AddUser(user)
	models.AddUser(user)

	for _, u := range models.GetUsers() {

		fmt.Printf("models.Users: %v\n", *u)
	}

	contollers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
