package game

import "github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"

type Engine struct {
}

type UserRepository interface {
	GetUser(userId string) (user_model.UserData, error)
}

type UserRepositoryImpl struct{}

func (repo *UserRepositoryImpl) GetUser(userId string) (user_model.UserData, error) {
	return user_model.UserData{
		U_seq: 10295032,
		U_id:  "1765733351097",
	}, nil
}
