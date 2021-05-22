package postgresDB

import (
	"2021_1_Noskool_team/internal/app/profiles"
	"2021_1_Noskool_team/internal/app/profiles/models"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
)

type ProfileRepository struct {
	con *sql.DB
}

func NewProfileRepository(con *sql.DB) profiles.Repository {
	return &ProfileRepository{
		con: con,
	}
}

// Create ...
func (r *ProfileRepository) Create(u *models.UserProfile) error {
	if err := u.ValidateForCreate(); err != nil {
		validationErr, _ := json.Marshal(err)
		return fmt.Errorf(string(validationErr))
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	_, err := r.con.Exec("INSERT INTO Profiles"+
		"(email, nickname, encrypted_password, avatar)"+
		"VALUES ($1, $2, $3, $4);",
		u.Email,
		u.Login,
		u.EncryptedPassword,
		u.Avatar)

	pgErr, ok := err.(*pq.Error)
	if ok {
		err = formattingDBerr(pgErr)
	}
	return err
}

// Update ...
func (r *ProfileRepository) Update(u *models.UserProfile) error {
	if err := u.ValidationForUpdate(); err != nil {
		fmt.Println(">>>", err)
		validationErr, _ := json.Marshal(err)
		return fmt.Errorf(string(validationErr))
	}

	_, err := r.con.Exec("UPDATE Profiles "+
		"SET email = $1, nickname = $2, first_name = $3, second_name = $4, encrypted_password = $5, favorite_genre = $6 "+
		"WHERE profiles_id = $7;",
		u.Email,
		u.Login,
		u.Name,
		u.Surname,
		u.EncryptedPassword,
		pq.Array(u.FavoriteGenre),
		u.ProfileID)

	pgErr, ok := err.(*pq.Error)
	if ok {
		err = formattingDBerr(pgErr)
	}
	return err
}

// FindByID ...
func (r *ProfileRepository) FindByID(id string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, first_name, second_name,"+
		"encrypted_password, avatar, favorite_genre  FROM Profiles"+
		" WHERE profiles_id = %s;", id)
	if err := r.con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.Name,
		&u.Surname,
		&u.EncryptedPassword,
		&u.Avatar,
		pq.Array(&u.FavoriteGenre)); err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateAvatar ...
func (r *ProfileRepository) UpdateAvatar(userID string, newAvatar string) {
	r.con.QueryRow("UPDATE Profiles "+
		"SET avatar = $1 WHERE profiles_id = $2;",
		newAvatar, userID)
}

// FindByLogin ...
func (r *ProfileRepository) FindByLogin(nickname string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, first_name, second_name, encrypted_password, avatar, favorite_genre  FROM Profiles"+
		" WHERE nickname = '%s';", nickname)
	if err := r.con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.Name,
		&u.Surname,
		&u.EncryptedPassword,
		&u.Avatar,
		pq.Array(&u.FavoriteGenre)); err != nil {
		return nil, err
	}
	return u, nil
}

// Create ...
func (r *ProfileRepository) UpdatePassword(id int, newPass string) error {
	if err := models.ValidateForChangePass(newPass); err != nil {
		return err
	}
	encNewPass, err := models.BeforeUpdatePass(newPass)
	if err != nil {
		return err
	}

	_, err = r.con.Exec("UPDATE Profiles SET encrypted_password = $1 WHERE profiles_id = $2;", encNewPass, id)

	if err != nil {
		return err
	}
	return nil
}

func formattingDBerr(err *pq.Error) error {
	fmt.Println(err)
	var formatedErr error
	switch {
	case err.Code == "23505" && err.Constraint == "profiles_email_key":
		formatedErr = models.ErrConstraintViolationEmail
	case err.Code == "23505" && err.Constraint == "profiles_nickname_key":
		formatedErr = models.ErrConstraintViolationNickname
	default:
		formatedErr = models.ErrDefaultDB
	}
	return formatedErr
}
