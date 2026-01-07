package game

import (
	"database/sql"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/KhoalaS/guitar-girl-offline/pkg/model/user_model"
)

const dateFormat string = "2006-01-02 15:04:05"

type Engine struct {
}

type UserRepository interface {
	GetUserByMemberId(memberId string) (user_model.UserData, error)
	GetUserBySequence(userSequence int32) (user_model.UserData, error)
	SaveUser(userSequence int32, data user_model.UserData) error
	CreateUser(memberId string, deviceUuid string) (user_model.UserData, error)
	UpdateLastCommunication(memberId string) error
}

type UserRepositoryImpl struct {
	database *sql.DB
}

func NewUserRepositoryImpl(database *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		database: database,
	}
}

func (repo *UserRepositoryImpl) GetUserByMemberId(memberId string) (user_model.UserData, error) {
	queryString := `SELECT * from user_data where uuid = ?`
	row := repo.database.QueryRow(queryString, memberId)
	var user user_model.UserData
	var uuid string

	err := row.Scan(
		&uuid,
		&user.U_seq,
		&user.U_id,
		&user.U_name,
		&user.U_nick,
		&user.U_cp,
		&user.U_candy,
		&user.U_like,
		&user.U_fans,
		&user.U_fans_grade,
		&user.U_selected_costume_id,
		&user.U_selected_music_id,
		&user.U_retain_continuous_tap,
		&user.U_join_type,
		&user.U_last_login,
		&user.U_last_communication,
		&user.U_save_date,
		&user.U_create_time,
		&user.U_tutorial_step,
		&user.U_review_popup,
		&user.Device_uuid,
		&user.U_samseck_step,
		&user.U_free_cp,
		&user.U_charge_cp,
	)
	if err != nil {
		return user_model.UserData{}, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) GetUserBySequence(userSequence int32) (user_model.UserData, error) {
	queryString := `SELECT * from user_data where u_seq = ?`
	row := repo.database.QueryRow(queryString, userSequence)

	var user user_model.UserData

	err := row.Scan(
		&user.U_seq,
		&user.U_id,
		&user.U_name,
		&user.U_nick,
		&user.U_cp,
		&user.U_candy,
		&user.U_like,
		&user.U_fans,
		&user.U_fans_grade,
		&user.U_selected_costume_id,
		&user.U_selected_music_id,
		&user.U_retain_continuous_tap,
		&user.U_join_type,
		&user.U_last_login,
		&user.U_last_communication,
		&user.U_save_date,
		&user.U_create_time,
		&user.U_tutorial_step,
		&user.U_review_popup,
		&user.Device_uuid,
		&user.U_samseck_step,
		&user.U_free_cp,
		&user.U_charge_cp,
	)

	if err != nil {
		return user_model.UserData{}, err
	}

	return user, nil
}

func (repo *UserRepositoryImpl) CreateUser(memberId string, deviceUuid string) (user_model.UserData, error) {
	var user user_model.UserData

	query := `
		INSERT INTO user_data (
			uuid,
			u_seq,
			u_id,
			u_name,
			u_nick,
			u_cp,
			u_candy,
			u_like,
			u_fans,
			u_fans_grade,
			u_selected_costume_id,
			u_selected_music_id,
			u_retain_continuous_tap,
			u_join_type,
			u_last_login,
			u_last_communication,
			u_save_date,
			u_create_time,
			u_tutorial_step,
			u_review_popup,
			device_uuid,
			u_samseck_step,
			u_free_cp,
			u_charge_cp
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`
	now := time.Now()

	uSeq := CreateUserSequence()
	uId := CreateUserId()
	loginDate := now.Format(dateFormat)

	_, err := repo.database.Exec(
		query,
		memberId,
		uSeq,                          // u_seq
		uId,                           // u_id
		"User",                        // u_name
		"User",                        // u_nick
		999999,                        // u_cp
		999999,                        // u_candy
		0,                             // u_like
		0,                             // u_fans
		0,                             // u_fans_grade
		1,                             // u_selected_costume_id
		1,                             // u_selected_music_id
		0,                             // u_retain_continuous_tap
		1,                             // u_join_type
		loginDate,                     // u_last_login
		"",                            // u_last_communication
		"",                            // u_save_date
		strconv.Itoa(int(now.Unix())), // u_create_time
		0,                             // u_tutorial_step
		"N",                           // u_review_popup
		deviceUuid,                    // device_uuid
		0,                             // u_samseck_step
		999999,                        // u_free_cp
		0,                             // u_charge_cp
	)
	if err != nil {
		return user, err
	}

	user.U_seq = uSeq
	user.U_id = uId

	return user, nil
}

func (repo *UserRepositoryImpl) SaveUser(userSequence int32, data user_model.UserData) error {
	// TODO
	return nil
}

func (repo *UserRepositoryImpl) UpdateLastCommunication(memberId string) error {
	nowDateString := time.Now().Format(dateFormat)

	query := `INSERT INTO user_data (u_last_communication) VALUES(?)`
	_, err := repo.database.Exec(query, nowDateString)

	if err != nil {
		return err
	}

	return nil
}

func CreateUserId() string {
	return strconv.Itoa(int(time.Now().UnixMilli()))
}

func CreateUserSequence() int32 {
	return rand.Int31n(math.MaxInt32)
}
