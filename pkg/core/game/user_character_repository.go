package game

import (
	"context"
	"database/sql"

	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
)

type UserCharacterRepository interface {
	GetCharacter(memberId string) ([]user_model.UserCharacter, error)
	SetCharacter(memberId string, characters []user_model.SaveUserCharacter) error
}

func NewUserCharacterRepository(database *sql.DB) UserCharacterRepository {
	return &UserCharacterRepositoryImpl{
		database: database,
	}
}

type UserCharacterRepositoryImpl struct {
	database *sql.DB
}

func (u *UserCharacterRepositoryImpl) GetCharacter(memberId string) ([]user_model.UserCharacter, error) {
	characters := []user_model.UserCharacter{}

	query := `SELECT i_id, i_level, i_bonusLevel FROM user_character WHERE uuid = ?`

	rows, err := u.database.Query(query, memberId)
	if err != nil {
		return characters, err
	}

	for rows.Next() {
		var character user_model.UserCharacter

		err := rows.Scan(
			&character.I_id,
			&character.I_Level,
			&character.I_BonusLevel,
		)
		if err != nil {
			rows.Close()
			return []user_model.UserCharacter{}, err
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func (u *UserCharacterRepositoryImpl) SetCharacter(memberId string, characters []user_model.SaveUserCharacter) error {
	ctx := context.TODO()
	tx, err := u.database.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
		INSERT INTO user_character VALUES (?,?,?,?)
		ON CONFLICT (uuid, i_id) DO
		UPDATE SET 
			i_level = excluded.i_level,
			i_bonusLevel = excluded.i_bonusLevel
	`

	for _, character := range characters {
		_, err := tx.Exec(query, memberId, character.I_id, character.I_Level, character.I_BonusLevel)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
