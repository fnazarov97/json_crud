package storage

import "github.com/fnazarov97/json_crud/models"

type StorageI interface {
	Close()
	User() UserRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	GetById(req *models.UserPrimaryKey) (*models.User, error)
	GetList(req *models.GetUserListReq) (*models.GetUserListRes, error)
}
