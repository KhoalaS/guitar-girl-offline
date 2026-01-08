package game

import (
	"database/sql"

	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
)

type UserAreaRepository interface {
	SetAreaData(memberId string, areaData user_model.SaveUserAreaInfo) error
	GetAreaData(memberId string, areaId int16) (user_model.UserAreaData, error)
	GetAllAreaData(memberId string) ([]user_model.UserAreaData, error)
}

type UserAreaRepositoryImpl struct {
	database *sql.DB
}

func (userAreaRepo *UserAreaRepositoryImpl) GetAllAreaData(memberId string) ([]user_model.UserAreaData, error) {
	areaDataList := []user_model.UserAreaData{}

	query := `SELECT * FROM user_area_info WHERE uuid = ?`
	rows, err := userAreaRepo.database.Query(query, memberId)

	if err != nil {
		return areaDataList, err
	}

	var _uuid string

	for rows.Next() {
		var areaData user_model.UserAreaData
		err := rows.Scan(
			&_uuid,
			&areaData.U_area_num,
			&areaData.D_Candy,
			&areaData.D_Like,
			&areaData.I_UserFanCount,
			&areaData.I_UserFanGrade,
			&areaData.I_SelectedCostumeId,
			&areaData.I_SelectedMusicId,
			&areaData.I_SelectedGuitarId,
			&areaData.S_TutorialList,
			&areaData.S_Gp1,
			&areaData.S_Gp2,
		)

		if err != nil {
			rows.Close()
			return areaDataList, err
		}

		areaDataList = append(areaDataList, areaData)
	}

	return areaDataList, nil
}

func (userAreaRepo *UserAreaRepositoryImpl) GetAreaData(memberId string, areaId int16) (user_model.UserAreaData, error) {
	query := `SELECT * FROM user_area_info WHERE uuid = ? AND u_area_num = ?`
	row := userAreaRepo.database.QueryRow(query, memberId, areaId)

	var areaData user_model.UserAreaData
	var uuid string

	err := row.Scan(
		&uuid,
		&areaData.U_area_num,
		&areaData.D_Candy,
		&areaData.D_Like,
		&areaData.I_UserFanCount,
		&areaData.I_UserFanGrade,
		&areaData.I_SelectedCostumeId,
		&areaData.I_SelectedMusicId,
		&areaData.I_SelectedGuitarId,
		&areaData.S_TutorialList,
		&areaData.S_Gp1,
		&areaData.S_Gp2,
	)

	if err != nil {
		return user_model.UserAreaData{}, err
	}

	return areaData, nil
}

func (userAreaRepo *UserAreaRepositoryImpl) SetAreaData(memberId string, areaData user_model.SaveUserAreaInfo) error {
	upsertQuery := `
	INSERT INTO user_area_info
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT (uuid, u_area_num) DO
		UPDATE SET 
			d_like = excluded.d_like,
    		i_user_fan_count = excluded.i_user_fan_count,
    		i_user_fan_grade = excluded.i_user_fan_grade,
    		i_selected_costume_id = excluded.i_selected_costume_id,
    		i_selected_music_id = excluded.i_selected_music_id,
    		i_selected_guitar_id = excluded.i_selected_guitar_id,
    		d_candy = excluded.d_candy,
    		s_tutorial_list = excluded.s_tutorial_list,
    		s_gp1 = excluded.s_gp1,
    		s_gp2 = excluded.s_gp2;`

	_, err := userAreaRepo.database.Exec(upsertQuery,
		memberId,
		areaData.U_area_num,
		areaData.D_Candy,
		areaData.D_Like,
		areaData.I_UserFanCount,
		areaData.I_UserFanGrade,
		areaData.I_SelectedCostumeId,
		areaData.I_SelectedMusicId,
		areaData.I_SelectedGuitarId,
		areaData.S_TutorialList,
		areaData.S_Gp1,
		areaData.S_Gp2,
	)

	if err != nil {
		return err
	}

	return nil

}

func NewUserAreaRepositoryImpl(database *sql.DB) UserAreaRepository {
	return &UserAreaRepositoryImpl{
		database: database,
	}
}
