package main

import (
	"fmt"
	"log"

	"github.com/fnazarov97/json_crud/config"
	"github.com/fnazarov97/json_crud/controller"
	"github.com/fnazarov97/json_crud/models"
	"github.com/fnazarov97/json_crud/storage/jsondb"
)

func main() {
	cfg := config.Load()

	jsondb, err := jsondb.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file" + err.Error())
	}
	defer jsondb.Close()

	c := controller.NewController(&cfg, jsondb)

	// id, err := c.Create(&models.CreateUser{
	// 	Name:    "Farhod",
	// 	Surname: "Nazarov",
	// 	Age:     26,
	// })

	// if err != nil {
	// 	log.Panicln("error whiling create user")
	// }
	// user, err := c.GetById(&models.UserPrimaryKey{Id: "f7cbfb94-fba3-4bbb-bcc6-373da581467b"})
	// if err != nil {
	// 	log.Panicln("error whiling GetById user", err)
	// }
	// fmt.Println(*user)
	users, err := c.GetList(&models.GetUserListReq{Offset: 0, Limit: 0, Search: ""})
	if err != nil {
		log.Panicln("error whiling GetById user", err)
	}
	fmt.Println(users.Cout, users.Users[0])

}
