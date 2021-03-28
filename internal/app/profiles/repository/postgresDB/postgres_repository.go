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
	if err := u.Validate(true); err != nil {
		validationErr, _ := json.Marshal(err)
		return fmt.Errorf(string(validationErr))
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	sqlReq := fmt.Sprintf("INSERT INTO Profiles"+
		"(email, nickname, first_name, second_name, encrypted_password, avatar, favorite_genre)"+
		"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s');",
		u.Email,
		u.Login,
		u.Name,
		u.Surname,
		u.EncryptedPassword,
		u.Avatar,
		u.FavoriteGenre)
	_, err := r.con.Exec(sqlReq)
	pgErr, ok := err.(*pq.Error)
	if ok {
		err = formattingDBerr(pgErr)
	}
	return err
}

// Update ...
func (r *ProfileRepository) Update(u *models.UserProfile, withPassword bool) error {
	if withPassword {
		if err := u.Validate(true); err != nil {
			validationErr, _ := json.Marshal(err)
			return fmt.Errorf(string(validationErr))
		}
		if err := u.BeforeCreate(); err != nil {
			return err
		}
	}
	if err := u.Validate(false); err != nil {
		validationErr, _ := json.Marshal(err)
		return fmt.Errorf(string(validationErr))
	}

	sqlReq := fmt.Sprintf("UPDATE Profiles "+
		"SET email = '%s', nickname = '%s', first_name = '%s', second_name = '%s', encrypted_password = '%s', favorite_genre = '%s' "+
		"WHERE profiles_id = '%d';",
		u.Email,
		u.Login,
		u.Name,
		u.Surname,
		u.EncryptedPassword,
		u.FavoriteGenre,
		u.ProfileID)
	_, err := r.con.Exec(sqlReq)
	pgErr, ok := err.(*pq.Error)
	if ok {
		err = formattingDBerr(pgErr)
	}
	return err
}

// FindByID ...
func (r *ProfileRepository) FindByID(id string) (*models.UserProfile, error) {
	u := &models.UserProfile{}
	sqlReq := fmt.Sprintf("SELECT profiles_id, email, nickname, first_name, second_name, encrypted_password, avatar, favorite_genre  FROM Profiles"+
		" WHERE profiles_id = %s;", id)
	if err := r.con.QueryRow(sqlReq).Scan(
		&u.ProfileID,
		&u.Email,
		&u.Login,
		&u.Name,
		&u.Surname,
		&u.EncryptedPassword,
		&u.Avatar,
		&u.FavoriteGenre); err != nil {
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
		&u.FavoriteGenre); err != nil {
		return nil, err
	}
	return u, nil
}

func formattingDBerr(err *pq.Error) error {
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
