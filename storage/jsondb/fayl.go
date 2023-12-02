package jsondb

import (
	"os"

	"github.com/fnazarov97/json_crud/config"
	"github.com/fnazarov97/json_crud/storage"
)

type Store struct {
	user *userRepo
}

func NewFileJson(cfg *config.Config) (storage.StorageI, error) {
	userFile, err := os.Open(cfg.Path + cfg.UserFileName)

	if err != nil {
		return nil, err
	}

	return &Store{
		user: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
	}, nil
}

func (s *Store) Close() {
	s.user.file.Close()
}

func (s *Store) User() storage.UserRepoI {
	return s.user
}
