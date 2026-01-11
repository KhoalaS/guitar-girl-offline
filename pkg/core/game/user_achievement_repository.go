package game

import (
	"context"
	"database/sql"

	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
)

type UserAchievementRepository interface {
	GetAchievements(memberId string) ([]user_model.UserAchievement, error)
	SetAchievements(memberId string, achievements []user_model.UserAchievement) error
}

func NewUserAchievementRepository(database *sql.DB) UserAchievementRepository {
	return &UserAchievementRepositoryImpl{
		database: database,
	}
}

type UserAchievementRepositoryImpl struct {
	database *sql.DB
}

func (u *UserAchievementRepositoryImpl) GetAchievements(memberId string) ([]user_model.UserAchievement, error) {
	achievements := []user_model.UserAchievement{}

	query := `SELECT (i_id, i_level, d_quantity, s_quantity) FROM user_achievement WHERE uuid = ?`
	rows, err := u.database.Query(query, memberId)
	if err != nil {
		return achievements, err
	}

	for rows.Next() {
		var achievement user_model.UserAchievement
		err := rows.Scan(
			&achievement.I_id,
			&achievement.I_Level,
			&achievement.D_Quantity,
			&achievement.S_Quantity,
		)

		if err != nil {
			rows.Close()
			return []user_model.UserAchievement{}, err
		}
	}

	return achievements, nil

}

func (u *UserAchievementRepositoryImpl) SetAchievements(memberId string, achievements []user_model.UserAchievement) error {
	ctx := context.TODO()
	tx, err := u.database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
		INSERT INTO user_achievements VALUES (?,?,?,?,?)
		ON CONFLICT (uuid, i_id) DO
		UPDATE SET 
			i_level = excluded.i_level,
			d_quantity = excluded.d_quantity,
			s_quantity = excluded.s_quantity
	`

	for _, achievement := range achievements {
		_, err := tx.Exec(query, memberId, achievement.I_id, achievement.I_Level, achievement.D_Quantity, achievement.S_Quantity)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
