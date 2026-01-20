package game

import (
	"context"
	"database/sql"

	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
)

type UserCostumeRepository interface {
	GetCostumes(memberId string) ([]user_model.UserCostume, error)
	SetCostumes(memberId string, costumes []user_model.SaveUserCostume) error
}

type UserCostumeRepositoryImpl struct {
	database *sql.DB
}

func NewUserCostumeRepository(database *sql.DB) UserCostumeRepository {
	return &UserCostumeRepositoryImpl{
		database: database,
	}
}

func (u *UserCostumeRepositoryImpl) GetCostumes(memberId string) ([]user_model.UserCostume, error) {
	costumes := []user_model.UserCostume{}

	query := `SELECT i_id, i_Level, i_BonusLevel from user_costume WHERE uuid = ?`
	rows, err := u.database.Query(query)
	if err != nil {
		return costumes, err
	}

	for rows.Next() {
		var costume user_model.UserCostume

		err := rows.Scan(&costume.I_id, &costume.I_Level, &costume.I_BonusLevel)
		if err != nil {
			rows.Close()
			return costumes, err
		}

		costumes = append(costumes, costume)
	}

	return costumes, nil
}

func (u *UserCostumeRepositoryImpl) SetCostumes(memberId string, costumes []user_model.SaveUserCostume) error {
	ctx := context.TODO()

	tx, err := u.database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO user_costume VALUES(?,?,?,?)
		ON CONFLICT (uuid, i_id) DO
		UPDATE SET 
			i_Level = excluded.i_Level,
			i_BonusLevel = excluded.i_BonusLevel

		`
	for _, costume := range costumes {
		_, err := u.database.Exec(query, memberId, costume.I_id, costume.I_Level, costume.I_BonusLevel)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
