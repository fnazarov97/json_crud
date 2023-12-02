package jsondb

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/fnazarov97/json_crud/models"

	"github.com/google/uuid"
)

type userRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewUserRepo(faylName string, file *os.File) *userRepo {
	return &userRepo{
		fileName: faylName,
		file:     file,
	}
}

func (u *userRepo) Create(req *models.CreateUser) (id string, err error) {
	id = uuid.New().String()
	var users = []*models.User{}
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return "", err
	}
	users = append(users, &models.User{
		Id:      id,
		Name:    req.Name,
		Surname: req.Surname,
		Age:     req.Age,
	})
	usersJson, err := json.MarshalIndent(users, "    ", "  ")
	if err != nil {
		return "", err
	}
	os.WriteFile(u.fileName, usersJson, os.ModeAppend)
	return id, nil
}

func (u *userRepo) GetById(req *models.UserPrimaryKey) (user *models.User, err error) {
	var users = []*models.User{}
	file, err := os.Open(u.fileName)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(byteValue, &users)
	// err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return &models.User{}, err
	}
	for _, u := range users {
		if u.Id == req.Id {
			user = u
		}
	}
	if user == nil {
		return &models.User{}, errors.New("not found")
	}
	return user, nil
}

func (u *userRepo) GetList(req *models.GetUserListReq) (res *models.GetUserListRes, err error) {
	var users = []*models.User{}
	file, err := os.Open(u.fileName)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil, err
	}
	return &models.GetUserListRes{Cout: len(users), Users: users}, nil
}
