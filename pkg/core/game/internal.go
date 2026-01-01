package game

import "github.com/KhoalaS/guitar-girl-offline/pkg/model"

type Engine struct {
}

type UserRepository interface {
	GetUser(userId string) (model.UserData, error)
}

type UserRepositoryImpl struct{}

func (repo *UserRepositoryImpl) GetUser(userId string) (model.UserData, error) {
	return model.UserData{
		U_seq: 10295032,
		U_mp:  "1765733351097",
	}, nil
}
