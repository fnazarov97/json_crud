package controller

import (
	"github.com/fnazarov97/json_crud/models"
)

func (c *Controller) Create(req *models.CreateUser) (id string, err error) {
	id, err = c.store.User().Create(req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Controller) GetById(req *models.UserPrimaryKey) (user *models.User, err error) {
	user, err = c.store.User().GetById(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Controller) GetList(req *models.GetUserListReq) (users *models.GetUserListRes, err error) {
	users, err = c.store.User().GetList(req)
	if err != nil {
		return nil, err
	}
	return users, nil
}
